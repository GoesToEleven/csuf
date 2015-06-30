package bulletinboard

import (
	"net/http"
	"text/template"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

type Post struct {
	Author     string
	Message    string
	UpdateDate string
	PostDate   string
	OP         bool
}

type Thread struct {
	Post Post
	ID   string
}

func init() {
	http.Handle("/hiddenDir/", http.StripPrefix("/hiddenDir/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", board)
	http.HandleFunc("/viewthread", viewthread)
	http.HandleFunc("/view", view)
	http.HandleFunc("/edit", editpost)
	http.HandleFunc("/create", createthread)
	http.HandleFunc("/delete", deletepost)
	http.HandleFunc("/post", createpost)
	http.HandleFunc("/put", put)
}

// The following handlerfunc was my first option for login and authentication.
// I decided to go with the authentication in the app.yaml.
/*func board(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-type", "text/html")
	c := appengine.NewContext(req)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, req.URL.String())
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Location", url)
		rw.WriteHeader(http.StatusFound)
		return
	}
	t := template.Must(template.ParseFiles("assets/board.html"))
	t.ExecuteTemplate(wr, "Board", nil)
}*/

func board(rw http.ResponseWriter, req *http.Request) {
	// rw.Header().Set("Content-type", "text/html")
	c := appengine.NewContext(req)
	u := user.Current(c)
	// rw.Header().Set("Location", req.URL.String())
	// rw.WriteHeader(http.StatusFound)

	posts := []Post{}
	q := datastore.NewQuery("post").Filter("OP =", true).Order("-PostDate")
	k, err := q.GetAll(c, &posts)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	threads := []Thread{}
	for i, value := range posts {
		threads = append(threads, Thread{value, k[i].Encode()})
	}

	t := template.Must(template.ParseFiles("assets/board.html"))
	t.ExecuteTemplate(rw, "Board", struct {
		User   string
		Thread []Thread
	}{
		u.String(),
		threads,
	})
}

func viewthread(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/view", http.StatusTemporaryRedirect)
}

func view(rw http.ResponseWriter, req *http.Request) {
	s := req.FormValue("encoded_key")
	c := appengine.NewContext(req)
	k, err := datastore.DecodeKey(s)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	posts := []Post{}
	q := datastore.NewQuery("post").Ancestor(k).Order("PostDate")
	keys, err := q.GetAll(c, &posts)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	threads := []Thread{}

	for i, value := range posts {
		threads = append(threads, Thread{value, keys[i].Encode()})
	}

	t := template.Must(template.ParseFiles("assets/post.html"))
	t.ExecuteTemplate(rw, "Post", struct {
		Parent string
		Thread []Thread
	}{
		s,
		threads,
	})
}

func editpost(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	s := req.FormValue("encoded_key")
	k, err := datastore.DecodeKey(s)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	mypost := Post{}
	err = datastore.Get(c, k, &mypost)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if mypost.Author == user.Current(c).String() {
		message := mypost.Message
		title := "Edit a Post"
		t := template.Must(template.ParseFiles("assets/edit.html"))
		t.ExecuteTemplate(rw, "New", struct {
			Title   string
			Message string
			ID      string
			Parent  string
		}{Title: title, Message: message, ID: s})
	} else {
		http.Redirect(rw, req, "/", http.StatusOK)
	}
}

func deletepost(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	u := user.Current(c)
	s := req.FormValue("encoded_key")
	k, err := datastore.DecodeKey(s)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	mypost := Post{}
	err = datastore.Get(c, k, &mypost)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if mypost.Author == u.String() {
		if mypost.OP {
			q := datastore.NewQuery("post").Ancestor(k).KeysOnly()
			keys, err := q.GetAll(c, struct{}{})
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			err = datastore.DeleteMulti(c, keys)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			err := datastore.Delete(c, k)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		// 	http.Redirect(rw, req, "/view", http.StatusTemporaryRedirect)
		// } else {
		http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
	}
}

func createthread(rw http.ResponseWriter, req *http.Request) {
	title := "Create a Thread"
	t := template.Must(template.ParseFiles("assets/edit.html"))
	t.ExecuteTemplate(rw, "New", struct {
		Title   string
		Message string
		ID      string
		Parent  string
	}{Title: title})
}

func createpost(rw http.ResponseWriter, req *http.Request) {
	title := "Create a Post"
	t := template.Must(template.ParseFiles("assets/edit.html"))
	t.ExecuteTemplate(rw, "New", struct {
		Title   string
		Message string
		ID      string
		Parent  string
	}{Title: title, Parent: req.FormValue("parent_key")})
}

func put(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	u := user.Current(c)
	m := req.FormValue("message")
	s := req.FormValue("encoded_key")
	// fmt.Fprintf(rw, "Key 1: %v", s)
	p := req.FormValue("parent_key")
	var t, ut string
	var op bool
	var k *datastore.Key

	// make/decode keys
	if s == "" {
		if p == "" {
			k = datastore.NewIncompleteKey(c, "post", nil)
			op = true
		} else {
			pk, err := datastore.DecodeKey(p)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			k = datastore.NewIncompleteKey(c, "post", pk)
			op = false
		}
		t = time.Now().Format("Jan 2, 2006 3:04 PM")
		ut = ""
	} else {
		k, err := datastore.DecodeKey(s)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		mypost := Post{}
		err = datastore.Get(c, k, &mypost)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		ut = time.Now().Format("Jan 2, 2006 3:04 PM")
		t = mypost.PostDate
		op = mypost.OP
	}

	// data := url.Values{}
	// data.Set("encoded_key", k.Encode())

	// r, _ := http.NewRequest("POST", "/view", bytes.NewBufferString(data.Encode()))

	newpost := Post{Author: u.String(), Message: m, UpdateDate: ut, PostDate: t, OP: op}
	_, err := datastore.Put(c, k, &newpost)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// http.Redirect(rw, r, "/view", http.StatusOK)
	http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
}
