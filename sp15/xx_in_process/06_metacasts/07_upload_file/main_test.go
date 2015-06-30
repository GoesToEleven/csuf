package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
	assert := assert.New(t)
	path := "/Users/markbates/Desktop/test_images/11.jpg"
	file, _ := os.Open(path)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("my_file", filepath.Base(path))
	io.Copy(part, file)
	writer.Close()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	dirname := "/tmp/go_file_uploads/TestUpload"
	os.RemoveAll(dirname)
	uploadDirectoryName = func() string {
		return dirname
	}

	Upload(res, req)

	assert.Equal(res.Code, 200)
	upload_file_path := fmt.Sprintf("%s/%s", dirname, filepath.Base(path))
	_, err := os.Stat(upload_file_path)
	assert.NoError(err)

	orig, _ := os.Open(path)
	defer orig.Close()
	orig_bytes := []byte{}
	orig.Read(orig_bytes)

	uploaded, _ := os.Open(upload_file_path)
	defer uploaded.Close()
	uploaded_bytes := []byte{}
	uploaded.Read(uploaded_bytes)

	assert.Equal(string(uploaded_bytes), string(orig_bytes))

	os.RemoveAll(dirname)
}
