package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/satori/go.uuid"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

type HTTPError struct {
	Code    int
	Message string
}

func (err HTTPError) Error() string {
	return err.Message
}

// READ ABOUT INTERFACE CONVERSIONS LIKE THIS: if err, ok := result.(HTTPError)
// https://golang.org/doc/effective_go.html#interface_conversions

func serveAPI(res http.ResponseWriter, req *http.Request, handler func() interface{}) {
	res.Header().Set("Content-Type", "application/json")

	result := handler()
	if err, ok := result.(HTTPError); ok {
		res.WriteHeader(err.Code)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Message,
		})
	} else if err, ok := result.(error); ok {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(map[string]string{
			"error": err.Error(),
		})
	} else if rc, ok := result.(io.ReadCloser); ok {
		io.Copy(res, rc)
		rc.Close()
	} else {
		json.NewEncoder(res).Encode(result)
	}
}

// Documents

type (
	DocumentFile struct {
		ID, Name string
	}
	Document struct {
		ID       string
		Link     string
		Contents string
		Files    []DocumentFile
	}
)

func serveDocumentsGet(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")
		email, ok := session.Values["email"].(string)
		if !ok {
			return HTTPError{403, "access denied"}
		}
		userKey := datastore.NewKey(ctx, "User", email, 0, nil)
		docKey := datastore.NewKey(ctx, "Document", params.ByName("id"), 0, userKey)
		var document Document
		err := datastore.Get(ctx, docKey, &document)
		if err != nil {
			return err
		}
		return document
	})
}

func serveDocumentsList(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")
		email, ok := session.Values["email"].(string)
		if !ok {
			return HTTPError{403, "access denied"}
		}
		userKey := datastore.NewKey(ctx, "User", email, 0, nil)
		documents := []Document{}
		_, err := datastore.NewQuery("Document").Ancestor(userKey).GetAll(ctx, &documents)
		if err != nil {
			return err
		}
		return documents
	})
}

func serveDocumentsCreate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")

		email, ok := session.Values["email"].(string)
		if !ok {
			return HTTPError{403, "access denied"}
		}

		var document Document
		err := json.NewDecoder(req.Body).Decode(&document)
		if err != nil {
			return err
		}
		if document.ID != "" {
			return fmt.Errorf("invalid document: id must not be set")
		}
		document.ID = uuid.NewV1().String()
		userKey := datastore.NewKey(ctx, "User", email, 0, nil)
		docKey := datastore.NewKey(ctx, "Document", document.ID, 0, userKey)
		docKey, err = datastore.Put(ctx, docKey, &document)
		if err != nil {
			return err
		}
		return document
	})
}

func serveDocumentsDelete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")

		email, ok := session.Values["email"].(string)
		if !ok {
			return HTTPError{403, "access denied"}
		}

		id := params.ByName("id")
		userKey := datastore.NewKey(ctx, "User", email, 0, nil)
		docKey := datastore.NewKey(ctx, "Document", id, 0, userKey)
		var document Document
		err := datastore.Get(ctx, docKey, &document)
		if err != nil {
			return err
		}
		err = datastore.Delete(ctx, docKey)
		if err != nil {
			return err
		}
		return true
	})
}

func serveDocumentsUpdate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")
		email, ok := session.Values["email"].(string)
		if !ok {
			return HTTPError{403, "access denied"}
		}

		var document Document
		err := json.NewDecoder(req.Body).Decode(&document)
		if err != nil {
			return err
		}
		document.ID = params.ByName("id")

		userKey := datastore.NewKey(ctx, "User", email, 0, nil)
		docKey := datastore.NewKey(ctx, "Document", params.ByName("id"), 0, userKey)
		var originalDocument Document
		err = datastore.Get(ctx, docKey, &originalDocument)
		if err != nil {
			return err
		}

		_, err = datastore.Put(ctx, docKey, &document)
		if err != nil {
			return err
		}

		return document
	})
}

// Files
func serveFilesGet(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")

		_, ok := session.Values["email"].(string)
		if !ok {
			return HTTPError{403, "access denied"}
		}

		bucket, err := file.DefaultBucketName(ctx)
		if err != nil {
			return err
		}

		hc := &http.Client{
			Transport: &oauth2.Transport{
				Source: google.AppEngineTokenSource(ctx, storage.ScopeReadOnly),
				Base:   &urlfetch.Transport{Context: ctx},
			},
		}

		cctx := cloud.NewContext(appengine.AppID(ctx), hc)
		rc, err := storage.NewReader(cctx, bucket, params.ByName("id"))
		if err != nil {
			return err
		}

		name := req.URL.Query().Get("name")
		if name == "" {
			name = params.ByName("id")
		}
		name = regexp.MustCompile("[^a-zA-Z-_.]").ReplaceAllString(name, "")

		res.Header().Set("Content-Disposition", "inline; filename=\""+name+"\"")
		res.Header().Set("Content-Type", "application/octet-stream")
		return rc
	})
}

func serveFilesUpload(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")

		_, ok := session.Values["email"].(string)
		if !ok {
			return HTTPError{403, "access denied"}
		}

		bucket, err := file.DefaultBucketName(ctx)
		if err != nil {
			return err
		}

		hc := &http.Client{
			Transport: &oauth2.Transport{
				Source: google.AppEngineTokenSource(ctx, storage.ScopeFullControl),
				Base:   &urlfetch.Transport{Context: ctx},
			},
		}

		id := uuid.NewV1().String()

		ff, _, err := req.FormFile("file")
		if err != nil {
			return err
		}
		defer ff.Close()

		cctx := cloud.NewContext(appengine.AppID(ctx), hc)
		wc := storage.NewWriter(cctx, bucket, id)
		io.Copy(wc, ff)
		err = wc.Close()
		if err != nil {
			return err
		}

		return id
	})
}

// Users

func serveUsersCreate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")

		type Request struct {
			Email, Password string
		}
		var request Request
		json.NewDecoder(req.Body).Decode(&request)

		err := CreateUser(ctx, request.Email, request.Password)
		if err != nil {
			return HTTPError{403, err.Error()}
		}
		session.Values["email"] = request.Email
		err = session.Save(req, res)
		if err != nil {
			return err
		}

		return true
	})
}

func serveUsersLogin(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		ctx := appengine.NewContext(req)
		session, _ := sessionStore.Get(req, "session")

		type Request struct {
			Email, Password string
		}
		var request Request
		json.NewDecoder(req.Body).Decode(&request)

		_, err := AuthenticateUser(ctx, request.Email, request.Password)
		if err != nil {
			return HTTPError{403, err.Error()}
		}

		session.Values["email"] = request.Email
		err = session.Save(req, res)
		if err != nil {
			return err
		}

		return true
	})
}

func serveUsersLogout(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serveAPI(res, req, func() interface{} {
		session, _ := sessionStore.Get(req, "session")
		session.Values = nil
		err := session.Save(req, res)
		if err != nil {
			return err
		}
		return true
	})
}
