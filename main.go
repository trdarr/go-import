package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const port = "80"

func main() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pkg := strings.TrimPrefix(r.URL.Path, "/")
		if pkg == "" {
			return
		}

		log.Printf("Got request for <%s>.\n", pkg)

		t.Execute(w, struct {
			Pkg string
			URI string
		}{pkg, fmt.Sprintf("https://github.com/trdarr/%s", pkg)})
	})

	log.Printf("Listening on port %s.\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
