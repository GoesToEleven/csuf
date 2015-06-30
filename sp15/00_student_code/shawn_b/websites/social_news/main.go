package main

import (
	"fmt"
//	"bytes"
	"net/http"
	"io"
	"time"
	"strings"
	"strconv"
//	"os"
	"math/rand"
	"html/template"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"

)


type webUser struct {
	Uuid uint64
	Nickname string
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/create2", create2Handler)
//	http.HandleFunc("/read", readHandler)
//	http.HandleFunc("/update", updateHandler)
//	http.HandleFunc("/update2", updateHandler2)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/edit2", editHandler2)
	http.HandleFunc("/delete", deleteHandler)

}



func handler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)
	log.Infof(c, "Got a visitor to the front page!")	//keep log in the imports

	//Check if the request is before or after to create the right query
	//The GET requests for the stories will be based around the SubmitDateTime
	//Using "after" will return stories after a certain date from newest to oldest
	//Using "before" will return stories before a certain date from oldest to newest
	//Default is to use the latest 3 submissions

	afterDate := r.FormValue("after")
	beforeDate := r.FormValue("before")
	returnLimit := 3
	showPrevLink := false

	var q *datastore.Query


	if afterDate != "" {
		showPrevLink = true
		//Get the results in descending order for newest to oldest
		afterDate = strings.Replace(afterDate, "%20", " ", -1) //replace all %20 with " "
		ttime, err := time.Parse(DateTimeDatastoreFormat, afterDate)
		if err != nil {
			serveError(c,w,err)
			return
		}
		q = datastore.NewQuery(WebSubmissionEntityName).
			Filter("SubmitDateTime <", ttime ).
			Order("-SubmitDateTime").
			Limit(returnLimit)
	} else if beforeDate != "" {
		showPrevLink = true
		//Get the results is ascending order from oldest to newest
		beforeDate = strings.Replace(beforeDate, "%20", " ", -1) //replace all %20 with " "
		ttime, err := time.Parse(DateTimeDatastoreFormat, beforeDate)
		if err != nil {
			serveError(c,w,err)
			return
		}
		q = datastore.NewQuery(WebSubmissionEntityName).
			Filter("SubmitDateTime >", ttime ).
			Order("SubmitDateTime").
			Limit(returnLimit)

		//limit check at the beginning if less than the returnLimit redo from the beginning
		length, cerr := q.Count(c)
		if cerr != nil {
			serveError(c,w,cerr)
		}

		//TODO refactor to not duplicate the default query below
		if length < returnLimit {
			showPrevLink = false
			q = datastore.NewQuery(WebSubmissionEntityName).
			Order("-SubmitDateTime").
			Limit(returnLimit)
		}
	} else {
		q = datastore.NewQuery(WebSubmissionEntityName).
			Order("-SubmitDateTime").
			Limit(returnLimit)

	}

	//Populate the results struct and store the keys
	var pageCon PageContainer

	for t := q.Run(c); ; {
		var x WebSubmission
		key, err := t.Next(&x)
		if err == datastore.Done{
			break
		}
		if err != nil{
//			serveError(c,w,err)
			fmt.Fprintf(w, "nope %v", err.Error())
			return
		}
		if u == nil {
			pageCon.Stories = append(pageCon.Stories, StoryListData{x, key, false})
		} else {
			pageCon.Stories = append(pageCon.Stories, StoryListData{x, key, u.String() == x.SubmitBy})
		}
	}

	//if we filled up the page with results there are probably more, build the
	//next page link
	length, cerr := q.Count(c)
	if cerr != nil {
		serveError(c,w,cerr)
	}
	if  length == returnLimit {
		//get the submit datetime of the last story
		pageCon.AfterLink = pageCon.Stories[returnLimit - 1].Story.SubmitDateTime.Format(DateTimeDatastoreFormat)
	}

	//If it was a prev page press reverse the result array to get it back into chronological order
	if length >= 1 && beforeDate != "" {
		for i, j := 0, len(pageCon.Stories)-1; i < j; i, j = i+1, j-1 {
			pageCon.Stories[i], pageCon.Stories[j] = pageCon.Stories[j], pageCon.Stories[i]
		}
	}

	//prev page link
	//check the length because going forward you can have null data
	if showPrevLink && length >= 1{
		pageCon.BeforeLink = pageCon.Stories[0].Story.SubmitDateTime.Format(DateTimeDatastoreFormat)
	}

	//build and show the page
	page := template.Must(template.ParseFiles(
		"public/templates/_base.html",
		"public/templates/storylist.html",
	))

	if err := page.Execute(w, pageCon); err != nil {
		serveError(c,w,err)
		fmt.Fprintf(w, "\n%v\n%v",err.Error(), pageCon)
		return
	}

}

func createHandler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	http.ServeFile(w,r, "public/templates/create.html")
}

func create2Handler(w http.ResponseWriter, r *http.Request) {
	strTitle := r.FormValue("title")
	strLink := r.FormValue("link")
//	fmt.Fprint(w, "go title: ", strTitle)
//	fmt.Fprint(w, "link: ", strLink)

	//TODO error checking for the title and link

	c := appengine.NewContext(r)
	u := user.Current(c)

	//TODO error checking for user?

	//TODO check random thread against an already existing one
	rand.Seed(time.Now().UnixNano())
	newThreadId := rand.Int63()

	newSubmission := WebSubmission {
		Title: strTitle,
		Link: strLink,
		SubmitBy: u.String(),
		Thread: newThreadId,	//TODO create random thread id
		SubmitDateTime: time.Now(),
		SubmissionDesc: "",
		Score: 0,
	}


	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, WebSubmissionEntityName, nil), &newSubmission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w,r,"/", http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	//For now identify by the ThreadId since each submission will have a unique one
	strVal := r.FormValue("thread")

	//TODO update with YAML require login to not need to check the user status

	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	page := template.Must(template.ParseFiles(
		"public/templates/_base_edit.html",
		"public/templates/edit.html",
	))

	data := struct{ThreadId string}{strVal}

	if err := page.Execute(w, data); err != nil {
		serveError(c,w,err)
		fmt.Fprintf(w, "\n%v\n%v",err.Error(), data)
		return
	}
}

func editHandler2(w http.ResponseWriter, r *http.Request) {
	strLink := r.FormValue("link")
	strThread := r.FormValue("thread")

	c := appengine.NewContext(r)

	threadId,_ := strconv.Atoi(strThread)

	q := datastore.NewQuery(WebSubmissionEntityName).
		Filter("Thread =", int64(threadId))

	//TODO cleanup checks whatever for emtpy or otherwise other than 1 result
	t := q.Run(c)

	var x WebSubmission
	key, err := t.Next(&x)
	if err != nil{
		serveError(c,w,err)
		fmt.Fprintf(w, "nope %v", err.Error())
		return
	}

	//Update field
	x.Link = strLink


	//Store back
	_, err2 := datastore.Put(c, key, &x)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w,r,"/",http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	strThread := r.FormValue("thread")
	c := appengine.NewContext(r)
	threadId,_ := strconv.Atoi(strThread)

	q := datastore.NewQuery(WebSubmissionEntityName).
		Filter("Thread =", int64(threadId))

	//TODO cleanup checks whatever for emtpy or otherwise other than 1 result
	t := q.Run(c)

	var x WebSubmission
	key, err := t.Next(&x)
	if err != nil{
		serveError(c,w,err)
		fmt.Fprintf(w, "nope %v", err.Error())
		return
	}

	err2 := datastore.Delete(c, key)
	if err != nil {
		serveError(c,w,err2)
		return
	}

	http.Redirect(w,r,"/", http.StatusFound)
}

func serveError(c context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Internal Server Error")
	log.Errorf(c, "%v", err)
}
