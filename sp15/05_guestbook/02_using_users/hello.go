package hello

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", myHandler)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// [START new_context]
	c := appengine.NewContext(r)
	// [END new_context]
	// [START get_current_user]
	u := user.Current(c)
	// [END get_current_user]
	// [START if_user]
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
	w.Header().Set("LedZepSong", "Black Dog")
	w.WriteHeader(http.StatusFound)
	// [END if_user]
	// [START output]
	fmt.Fprintf(w, "String returns a displayable name for the user: %v <br>", u)
	fmt.Fprintf(w, "Email: %v <br>", u.Email)
	fmt.Fprintf(w, "AuthDomain: %v <br>", u.AuthDomain)
	fmt.Fprintf(w, "Admin: %v <br>", u.Admin)
	fmt.Fprintf(w, "ID: %v <br>", u.ID)
	fmt.Fprintf(w, "FederatedIdentity: %v <br>", u.FederatedIdentity)
	fmt.Fprintf(w, "FederatedProvider: %v <br>", u.FederatedProvider)

	// [END output]
}

/*
notes:
https://cloud.google.com/appengine/docs/go/reference#NewContext
https://cloud.google.com/appengine/docs/go/reference#Context
https://cloud.google.com/appengine/docs/go/users/reference

see this running here:
http://trial2-907.appspot.com/
--note:
---- if it's not running, I took it down
---- you can follow these instructions to run it locally
https://cloud.google.com/appengine/docs/go/gettingstarted/helloworld
---- you can follow these instructions to run on app engine
https://cloud.google.com/appengine/docs/go/gettingstarted/uploading
*/
