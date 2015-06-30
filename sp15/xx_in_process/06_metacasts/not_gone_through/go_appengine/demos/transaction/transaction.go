// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package transaction

// This app demonstrates the datastore transaction API.
//
// A bank account is initialized with 4 million dollars, and 10 concurrent
// withdrawals are made, each for slightly less than 1 million dollars.
//
// At most 4 out of 10 transactions can succeed. The remainder will fail
// either for business logic reasons (insufficient funds) or because of a
// conflict with a concurrent transaction.
//
// Most transactions will be attempted more than once, and some successful
// transactions will require more than one attempt.

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Not Found\n")
}

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Internal Server Error\n")
	c.Errorf("%v", err)
}

type BankAccount struct {
	Balance int
}

func withdraw(c appengine.Context, sc chan string, id string, amount, nAttempts int) error {
	b := BankAccount{}
	key := datastore.NewKey(c, "BankAccount", "", 1, nil)
	if err := datastore.Get(c, key, &b); err != nil {
		return err
	}
	sc <- fmt.Sprintf("%s: balance is $%07d  (attempt number %d)\n", id, b.Balance, nAttempts)

	time.Sleep(time.Duration(5+rand.Intn(15)) * time.Millisecond)

	b.Balance -= amount
	if b.Balance < 0 {
		return errors.New("insufficient funds")
	}
	if _, err := datastore.Put(c, key, &b); err != nil {
		return err
	}
	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		serve404(w)
		return
	}

	c := appengine.NewContext(r)
	b := BankAccount{4e6}
	key := datastore.NewKey(c, "BankAccount", "", 1, nil)
	if _, err := datastore.Put(c, key, &b); err != nil {
		serveError(c, w, err)
		return
	}

	sc := make(chan string)
	donec := make(chan bool)
	const N = 10
	for i := 0; i < N; i++ {
		go func(id string, amount int) {
			// Spread out the withdrawal requests.
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

			var nAttempts int
			err := datastore.RunInTransaction(c, func(c appengine.Context) error {
				nAttempts++
				return withdraw(c, sc, id, amount, nAttempts)
			}, nil)
			if err != nil {
				sc <- fmt.Sprintf("%s: error: %v\n", id, err)
			} else {
				sc <- fmt.Sprintf("%s: success\n", id)
			}
			donec <- true
		}(fmt.Sprintf("#%04d", 1<<uint(i)), 1e6-1<<uint(i))
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for nDone := 0; nDone < N; {
		select {
		case s := <-sc:
			io.WriteString(w, s)
		case <-donec:
			nDone++
		}
	}

	if err := datastore.Get(c, key, &b); err != nil {
		serveError(c, w, err)
		return
	}
	fmt.Fprintf(w, "\nFinal balance: $%d\nSuccessful withdrawals:\n", b.Balance)
	for i := 0; i < N; i++ {
		x := 1 << uint(i)
		if (b.Balance%1e6)&x != 0 {
			fmt.Fprintf(w, "#%04d\n", x)
		}
	}
}

func init() {
	http.HandleFunc("/", handle)
}
