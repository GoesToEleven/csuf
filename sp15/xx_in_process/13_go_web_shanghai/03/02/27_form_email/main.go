package main

import (
    "github.com/julienschmidt/httprouter"
    "html/template"
    "log"
    "net/http"
    "regexp"
    "strings"
    "fmt"
    "net/smtp"
)

func main() {
    router := httprouter.New()
    router.GET("/", index)
    router.POST("/", send)
    router.GET("/confirmation", confirmation)

    log.Println("Listening tortuga ...")
    http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    render(w, "templates/index.html", nil)
}

func send(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    msg := &Message{
        Email: r.FormValue("email"),
        Content: r.FormValue("content"),
    }

    if msg.Validate() == false {
        render(w, "templates/index.html", msg)
        return
    }

    if err := msg.Deliver(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func confirmation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    render(w, "templates/confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
    tmpl, err := template.ParseFiles(filename)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

type Message struct {
    Email   string
    Content string
    Errors  map[string]string
}

func (msg *Message) Validate() bool {
    msg.Errors = make(map[string]string)

    re := regexp.MustCompile(".+@.+\\..+")
    matched := re.Match([]byte(msg.Email))
    if matched == false {
        msg.Errors["Email"] = "Please enter a valid email address"
    }

    if strings.TrimSpace(msg.Content) == "" {
        msg.Errors["Content"] = "Please write a message"
    }

    return len(msg.Errors) == 0
}

func (msg *Message) Deliver() error {
    to := []string{"tuddleymc@gmail.com"}
    body := fmt.Sprintf("Reply-To: %v\r\nSubject: New Message\r\n%v", msg.Email, msg.Content)

    username := "toddmcleod@gmail.com"
    password := "..."
    auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

    return smtp.SendMail("smtp.gmail.com:587", auth, msg.Email, to, []byte(body))
}
