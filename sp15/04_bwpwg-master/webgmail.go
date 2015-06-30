package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func main() {
       	http.HandleFunc("/", root)
	http.HandleFunc("/mail", mail)
        fmt.Println("listening...")
        err := http.ListenAndServe(GetPort(), nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}

func GetPort() string {
        var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Enter your mail</title>
        <link rel="stylesheet" href="http://yui.yahooapis.com/pure/0.4.2/pure-min.css">
      </head>
      <body>
        <h2>Enter your Email</h2>
        <form action="/mail" method="post" accept-charset="utf-8" class="pure-form">
          To:<input type="text" name="to" />
          Subject:<input type="text" name="subject" />
          Message:<input type="text" name="msg" />
	  <input type="submit" value="... and check the mail!"  class="pure-button pure-button-primary" />
        </form>
        <div>
          <p><b>&copy; 2014 RubyLearning. All rights reserved.</b></p>
        </div>
      </body>
    </html>
`

const mailTemplateHTML = ` 
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <title>Email Result</title>
      <link rel="stylesheet" href="http://yui.yahooapis.com/pure/0.4.2/pure-min.css">
    </head>
    <body>
      <h2>Email Result</h2>
      <p>Your email has been sent to: </p>
      <pre>{{html .}}</pre>
      <p><a href="/">Start again!</a></p>
      <div>
        <p><b>&copy; 2014 RubyLearning. All rights reserved.</b></p>
      </div>
    </body>
  </html>
`

var mailTemplate = template.Must(template.New("mail").Parse(mailTemplateHTML))

func mail(w http.ResponseWriter, r *http.Request) {
        to := r.FormValue("to")
        subject := r.FormValue("subject")
        message := r.FormValue("msg")

        body := "To: " + to + "\r\nSubject: " +
          subject + "\r\n\r\n" + message
          
        auth := smtp.PlainAuth("", "satish.talim", "password", "smtp.gmail.com")
        err := smtp.SendMail("smtp.gmail.com:587", auth, "satish.talim@gmail.com", 
          []string{to},[]byte(body))
        if err != nil {
                log.Fatal("SendMail: ", err)
        } 
        
        err1 := mailTemplate.Execute(w, to)
        if err1 != nil {
	        http.Error(w, err1.Error(), http.StatusInternalServerError)
        }
}


