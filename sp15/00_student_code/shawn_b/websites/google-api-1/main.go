package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"html/template"
//	"os"
	"log"
	"io/ioutil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
//	"golang.org/x/net/context"
	gaeLog "google.golang.org/appengine/log"
//	gmail "google.golang.org/api/gmail/v1"

)

type ClientSecret struct {
	Web WebType `json: "web"`
}

type WebType struct{
	Auth_uri string `json: "auth_uri"`
	Client_secret string `json: "client_secret"`
	Token_uri string `json: "token_uri"`
	Client_email string `json: "client_email"`
	Redirect_uris []string `json: "redirect_uris"`
	Client_x509_cert_url string `json: "client_x509_cert_url"`
	Client_id string `json: "client_id"`
	Auth_provider_x509_cert_url string `json: "auth_provider_x509_cert_url"`
	Javascript_origins []string `json: "javascript_origins"`
}

type gmailProfile struct {
	EmailAddress string `json: "emailAddress"`
	MessagesTotal uint64 `json: "messagesTotal"`
	ThreadsTotal uint64 `json: "threadsTotal"`
	HistoryId string `json: "historyId"`
}

var conf = new(oauth2.Config)
//var gaeCtx = appengine.NewContext

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/g_start", g_start)
	http.HandleFunc("/oauth2callback", oauth2callback)
	http.HandleFunc("/formResult", formHandler)
}

func oauth2callback( w http.ResponseWriter, r *http.Request) {

	//TODO: validate FormValue("state")

	code :=  r.FormValue("code")

	c := appengine.NewContext(r)

	//Example log to Gae Log
//	gaeLog.Infof(c, "State Val: %s", r.FormValue("state"))

	tok, err := conf.Exchange(c, code)
	if err != nil {
		log.Fatal(err)
	}

//	client := conf.Client(oauth2.NoContext, tok)
	client := urlfetch.Client(c)
//	client.Get("...")

	gaeLog.Infof(c, "Token: %v", tok)
	gaeLog.Infof(c, "Access Token: %v", tok.AccessToken)
//	fmt.Fprint(w, "No Autographs!")

	fullUrl := fmt.Sprintf(
		"https://www.googleapis.com/gmail/v1/users/me/messages?fields=resultSizeEstimate&maxResults=10&includeSpamTrash=false&q=%%22is%%3Aunread%%22&access_token=%s",
		tok.AccessToken)

	gaeLog.Infof(c, "GET Request %v", fullUrl)
//	req, err2 := http.NewRequest("GET", fullUrl, nil)
//	if err2 != nil {
//		gaeLog.Infof(c, "request: %s", err2.Error())
//		log.Fatal("NewRequest: ", err)
//	}

//	resp, requestErr := client.Do(req)
	resp, requestErr := client.Get(fullUrl)
	if requestErr != nil {
		gaeLog.Infof(c, "client.Do err: %v", requestErr.Error())
		log.Fatal("Do: ", requestErr)
	}

	defer resp.Body.Close()

	body, dataReadErr := ioutil.ReadAll(resp.Body)
	if dataReadErr != nil {
		gaeLog.Infof(c, "dataReadErr: %v", dataReadErr.Error())
		log.Fatal("ReadAll: ", dataReadErr)
	}
	gaeLog.Infof(c, "data received: %v", body)

	res := make(map[string]interface{},0)
	json.Unmarshal(body, &res)
	gaeLog.Infof(c, "unmarshaled: ", res)
	var unReadMsg = res["resultSizeEstimate"]

	fmt.Fprint(w,"Your unread messages: ", unReadMsg)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Go Gmail</title>
        <link rel="stylesheet" href="/stylesheets/goview.css">
      </head>
      <body>
        <h2>Unread Messages!?</h2>
        <p>Cick on the button to see how many uread messages you have:</p>
        <form style="margin-left: 120px;" action="/g_start" accept-charset="utf-8">
          <input type="submit" value="Go!" />
        </form>
      </body>
    </html>
`

func g_start(w http.ResponseWriter, r *http.Request) {
	//Read the client_secret.json and parse the file so we can get our Google authorization
	file, err := ioutil.ReadFile("./config/client_secret.json")
	if err != nil {
		fmt.Println("client_secret.json:", err)
	}

	var clientSecret ClientSecret
	err = json.Unmarshal(file, &clientSecret)
	if err != nil {
		fmt.Println("client_secret.json unmarshall err:", )
	}

	//TODO: update with ConfigFromJson
	conf = &oauth2.Config{
		ClientID: clientSecret.Web.Client_id,
		ClientSecret: clientSecret.Web.Client_secret,
		RedirectURL: clientSecret.Web.Redirect_uris[0],
		Scopes: []string {
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}
	if err != nil {
		log.Fatal(err)
	}

	//Redirect the user to Google sign-in or authorization page
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)

	http.Redirect(w, r, url, http.StatusFound)
}


var resultFile, _ = ioutil.ReadFile("templates/result.html");
var resultHtmlTemplate = template.Must(template.New("result").Parse(string(resultFile)))

func formHandler(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("str")

	var err error

	if strEntered == "Shawn" {
		err = resultHtmlTemplate.Execute(w, "You spelled it correctly")
	} else {
		err = resultHtmlTemplate.Execute(w, "...you spelled it wrong")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

