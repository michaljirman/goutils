package handler

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://www.example.com/", nil)
	w:= httptest.NewRecorder()
	Handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != "<html><body>Hello!</body></html>"{
		t.Fail()
	}
}