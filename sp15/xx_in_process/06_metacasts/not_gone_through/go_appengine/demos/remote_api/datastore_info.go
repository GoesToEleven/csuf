package main

// Simple tool which lists the entity kinds, and a sample for each, for an
// app engine app.
//
// This tool can be invoked using the goapp tool bundled with the SDK.
// $ goapp run demos/remote_api/datastore_info.go \
//   -email admin@example.com \
//   -host my-app.appspot.com \
//   -password_file ~/.my_password

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/remote_api"
)

var (
	host         = flag.String("host", "", "hostname of application")
	email        = flag.String("email", "", "email of an admin user for the application")
	passwordFile = flag.String("password_file", "", "file which contains the user's password")
)

// See https://cloud.google.com/appengine/docs/go/datastore/stats
const DatastoreKindName = "__Stat_Kind__"

type DatastoreKind struct {
	KindName            string    `datastore:"kind_name"`
	EntityBytes         int       `datastore:"entity_bytes"`
	BuiltinIndexBytes   int       `datastore:"builtin_index_bytes"`
	BuiltinIndexCount   int       `datastore:"builtin_index_count"`
	CompositeIndexBytes int       `datastore:"composite_index_bytes"`
	CompositeIndexCount int       `datastore:"composite_index_count"`
	Timestamp           time.Time `datastore:"timestamp"`
	Count               int       `datastore:"count"`
	Bytes               int       `datastore:"bytes"`
}

func main() {
	flag.Parse()

	if *host == "" {
		log.Fatalf("Required flag: -host")
	}
	if *email == "" {
		log.Fatalf("Required flag: -email")
	}
	if *passwordFile == "" {
		log.Fatalf("Required flag: -password_file")
	}

	p, err := ioutil.ReadFile(*passwordFile)
	if err != nil {
		log.Fatalf("Unable to read password from %q: %v", *passwordFile, err)
	}
	password := strings.TrimSpace(string(p))

	client := clientLoginClient(*host, *email, password)

	c, err := remote_api.NewRemoteContext(*host, client)
	if err != nil {
		log.Fatalf("Failed to create context: %v", err)
	}
	log.Printf("App ID %q", appengine.AppID(c))

	q := datastore.NewQuery(DatastoreKindName).Order("kind_name")
	kinds := []*DatastoreKind{}
	if _, err := q.GetAll(c, &kinds); err != nil {
		log.Fatalf("Failed to fetch kind info: %v", err)
	}

	for _, k := range kinds {
		fmt.Printf("\nkind %q\t%d entries\t%d bytes\n", k.KindName, k.Count, k.Bytes)

		props := datastore.PropertyList{}
		if _, err := datastore.NewQuery(k.KindName).Limit(1).Run(c).Next(&props); err != nil {
			log.Printf("Unable to fetch sample entity kind %q: %v", k.KindName, err)
			continue
		}
		for _, prop := range props {
			fmt.Printf("\t%s: %v\n", prop.Name, prop.Value)
		}
	}
}

func clientLoginClient(host, email, password string) *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("failed to make cookie jar: %v", err)
	}
	client := &http.Client{
		Jar: jar,
	}

	v := url.Values{}
	v.Set("Email", email)
	v.Set("Passwd", password)
	v.Set("service", "ah")
	v.Set("source", "Misc-remote_api-0.1")
	v.Set("accountType", "HOSTED_OR_GOOGLE")

	resp, err := client.PostForm("https://www.google.com/accounts/ClientLogin", v)
	if err != nil {
		log.Fatalf("could not post login: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unsuccessful request: status %d; body %q", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatalf("unable to read response: %v", err)
	}

	m := regexp.MustCompile(`Auth=(\S+)`).FindSubmatch(body)
	if m == nil {
		log.Fatalf("no auth code in response %q", body)
	}
	auth := string(m[1])

	u := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     "/_ah/login",
		RawQuery: "continue=/&auth=" + url.QueryEscape(auth),
	}

	// Disallow redirects.
	redirectErr := errors.New("stopping redirect")
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return redirectErr
	}

	resp, err = client.Get(u.String())
	if urlErr, ok := err.(*url.Error); !ok || urlErr.Err != redirectErr {
		log.Fatalf("could not get auth cookies: %v", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusFound {
		log.Fatalf("unsuccessful request: status %d; body %q", resp.StatusCode, body)
	}

	client.CheckRedirect = nil
	return client
}
