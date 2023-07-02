// Package main runs a small server using embedded html files.
// The html files are all served in static context and the hello.html file is also served as a template
// The readme file is also served
package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed html
var htmlFS embed.FS

//go:embed README.md
var readme []byte // can also be embedded as a string

func main() {
	port := "8000"
	log.Println("starting server at http://localhost:" + port)
	http.ListenAndServe(":"+port, newHandler())
}

func newHandler() http.Handler {
	helloT := template.Must(template.New("hello.html").ParseFS(htmlFS, "html/hello.html"))
	mux := http.NewServeMux()
	mux.Handle("/hello", helloTemplateHandler(helloT))
	mux.Handle("/about", aboutHandler())
	mux.Handle("/", http.FileServer(http.FS(htmlFS)))
	return mux
}

func helloTemplateHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			name = "Gopher"
		}
		t.Execute(w, name)
	}
}

func aboutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(readme)
	}
}
