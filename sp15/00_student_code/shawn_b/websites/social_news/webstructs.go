package main

import (
	"time"
	"google.golang.org/appengine/datastore"
)

type WebSubmission struct {
	Title string
	Link string
	SubmitBy string
	Thread int64
	SubmitDateTime time.Time
	SubmissionDesc string
	Score int64
}

type PageContainer struct {
	Stories []StoryListData
	BeforeLink string
	AfterLink string
}

type StoryListData struct {
	Story WebSubmission
	Key *datastore.Key
	ShowEditDelete bool
}

const (
	WebSubmissionEntityName = "webSubmission"
	DateTimeDatastoreFormat = "2006-01-02 15:04:05.99 -0700 MST"	//The numbers used in this layout example matters! http://golang.org/src/time/format.go
	WebSubmissionEditName = "webSubmissionEdit"
)
