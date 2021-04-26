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
	h, err := newHandler()
	if err != nil {
		log.Fatalf("creating handler: %v", err)
	}
	svr := http.Server{
		Addr:    ":8000",
		Handler: h,
	}
	log.Println("starting server at http://localhost" + svr.Addr)
	svr.ListenAndServe()
}

func newHandler() (http.Handler, error) {
	helloFile, err := htmlFS.ReadFile("html/hello.html")
	if err != nil {
		return nil, err
	}
	t, err := template.New("").Parse(string(helloFile))
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.Handle("/hello", helloTemplateHandler(t))
	mux.Handle("/about", aboutHandler())
	mux.Handle("/", http.FileServer(http.FS(htmlFS)))
	return mux, nil
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
