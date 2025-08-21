package main

import (
	"fmt"

	"net/http"
	"text/template"
)

type Hit struct {
	URL   string // remember to use upper case for tmpl str
	Count float64
}

type TemplateData struct {
	Hits  []Hit
	Other int
	More  string
}

func Serve(nW *SearchWord, m *FullMap, count *CorpusCount, docs *Corpus) {
	//access handles for different parts of the server
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/top10", http.StripPrefix("/top10/", http.FileServer(http.Dir("./static/top10"))))
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		//getting the user given term
		word := r.URL.Query().Get("term")

		//change the word to the term
		nW.changeWord(word)

		//setting up style
		w.Header().Set("Content-Type", "text/html")

		//getting the amount of hits
		hits := TfIdf(nW, m, count, docs)

		//looks like the thing from lab03
		tmplBody := "<ol> {{range .}} <li>{{.URL}} {{.Count}}</li> {{end}} </ol>"
		tmpl, err := template.New("demo").Parse(tmplBody)
		if err != nil {
			fmt.Printf("template.Parse returned %v\n", err)
		}
		tmpl.Execute(w, hits)

	})
	http.ListenAndServe(":8080", nil)
}
