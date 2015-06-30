package main

import "github.com/gorilla/sessions"

var (
	// sessions are stored in a (tamper-proof) cookie
	sessionSecret = []byte("eUJE3c4Av3rAncC3yUSmdhjzHNuhbW8WuRfTdej8")
	sessionStore  *sessions.CookieStore
)

func init() {
	sessionStore = sessions.NewCookieStore(sessionSecret)
}
