package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
	"html/template"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	assets = flag.String("assets", ".", "path to assets")
	homeTempl *template.Template
)


// executes the template defined above
func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

// creates and runs a hub, parses the template, serves static assets on "/static", handles all the other stuff
func main() {
	flag.Parse()
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "template/index.html")))
	h := newHub()
	go h.run()
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", homeHandler)
	http.Handle("/ws", wsHandler{h: h})
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
