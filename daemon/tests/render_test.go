package test

import (
	"testing"
	"os"
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/astaxie/beego"
	"io/ioutil"
)

// TestRenderDocument test document rendering
func TestRenderDocument(t *testing.T) {
	wd, _ := os.Getwd()
	texFile := wd + "/" + "tests/g-brief2.test.tex"

	file, err := os.Open(texFile)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fi, err := file.Stat()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("document", fi.Name())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	req, err := http.NewRequest("POST", "/api/v1/document", body)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	assert.Equal(t, "application/pdf", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Header().Get("Content-Disposition"), "attachment; filename=\"")
	assert.Contains(t, w.Header().Get("Content-Disposition"), "-document.pdf\"")

	if len(w.Body.Bytes()) < 4 {
		t.Errorf("wrong body length")
		t.FailNow()
	}

	expectedMagicBytes := []byte("%PDF")
	respondedMagicBytes := w.Body.Bytes()[:4]

	assert.Equal(t, expectedMagicBytes, respondedMagicBytes)
}
