package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	handlerTests := []struct {
		url          string
		wantContents string
	}{
		{"/", "html/"}, // fileserver directories contains links to their contents
		{"/html/file1.html", "file1"},
		{"/html/file2.html", "file2"},
		{"/html/hello.html", "{{"},
		{"/hello", "Gopher"},
		{"/hello?name=fred", "fred"},
		{"/hello?name=Foo%20Bar", "Foo Bar"},
		{"/about", "go-server-embed-demo"},
		{"/INVALID", "404"}, // http.ServeMux has 404 a handler that prints the status code to the body
	}
	for _, test := range handlerTests {
		h, err := newHandler()
		if err != nil {
			t.Fatalf("unwanted error creating handler: %v", err)
		}
		r := httptest.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		got := w.Body.String()
		if !strings.Contains(got, test.wantContents) {
			t.Errorf("url '%v': wanted body to contain '%v', got '%v'", test.url, test.wantContents, got)
		}
	}
}
