package googleapi

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"log"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	// "io/ioutil"
	"net/http"
)

type resultinput struct {
	Search string
	Emails []string
}

func init() {
	http.Handle("/hiddenDir/", http.StripPrefix("/hiddenDir/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", apiQuery)
	http.HandleFunc("/result", results)
	http.HandleFunc("/process", process)
}

func apiQuery(rw http.ResponseWriter, req *http.Request) {
	url := `
	https://accounts.google.com/o/oauth2/auth?
	scope=https://www.googleapis.com/auth/gmail.readonly&
	redirect_uri=https%3A%2F%2Fcurious-cistern-90523.appspot.com%2Fprocess&,
	response_type=code&
	client_id=506452277892-llq9dpqoq9hu7h1ma13e6r7mahndil60.apps.googleusercontent.com&
	include_granted_scopes=true`

	t := template.Must(template.ParseFiles("assets/apiQuery.html"))
	err := t.ExecuteTemplate(rw, "apiQuery", url)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func process(rw http.ResponseWriter, req *http.Request) {
	err := req.URL.Query().Get("error")
	if err != "" {
		log.Fatalf("Unable to authenticate: %v", err)
	}

	code := req.URL.Query().Get("code")
	if code == "" {
		log.Fatal("Authentication Code Error")
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, body)
}

func results(rw http.ResponseWriter, req *http.Request) {
	var mysearch = resultinput{Search: req.FormValue("search")}

	// safeAddr := url.QueryEscape(mysearch.Search)
	// fullURL := fmt.Sprintf("https://www.googleapis.com/gmail/v1/users/me/messages?q=%s", mysearch)
	c := appengine.NewContext(req)
	// u, err := user.CurrentOAuth(c, "https://www.googleapis.com/auth/gmail.readonly")
	// if err != nil {
	// 	log.Fatalf("Unable to authenticate user: %v", err)
	// }

	client := urlfetch.Client(c)

	svc, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to create gmail service: %v", err)
	}

	res := svc.Users.Messages.List("me").Q(mysearch.Search)

	r, err := res.Do()

	// req, err := http.NewRequest("GET", fullURL, nil)
	// if err != nil {
	// 	log.Fatal("NewRequest: ", err)
	// }

	// resp, requestErr := client.Do(req)
	/*if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}

	// defer resp.Body.Close()

	// body, dataReadErr := ioutil.ReadAll(resp.Body)
	// if dataReadErr != nil {
	// 	log.Fatal("ReadAll: ", dataReadErr)
	// }

	// res := make(map[string][]map[string]interface{}, 0)

	// json.Unmarshal(body, &res)

	// for i, _ := range res["messages"] {
	// 	temp, _ = res["messages"][i]["id"]
	// 	append(mysearch.Emails, temp)
	// }

	// var myemails []string

	/*for _, value := range r.Messages {
		mysearch.Emails = append(mysearch.Emails, "https://mail.google.com/mail/u/0/#all/"+value.Id)
	}*/

	// t := template.Must(template.ParseFiles("assets/results.html"))
	// err = t.ExecuteTemplate(rw, "results", mysearch)
	// if err != nil {
	fmt.Fprintf(rw, "An error has occured again: SEARCH:%v CLIENT:%v SVC:%v RES:%v R:%v ERR1:%v", mysearch.Search, client, svc, res, r, err)
	// http.Error(rw, err.Error(), http.StatusInternalServerError)

	// }
}
