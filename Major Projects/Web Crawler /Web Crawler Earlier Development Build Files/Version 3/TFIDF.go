package main

import (
	"math"
	"slices"
	"sort"

	"github.com/kljensen/snowball"
)

type ByHits []Hit

func (p ByHits) Len() int {
	return len(p)
}

func (p ByHits) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByHits) Less(i, j int) bool {
	if p[i].Count == p[j].Count {
		return p[i].Count > p[j].Count
	}
	return p[i].Count < p[j].Count

}

func TfIdf(nW *SearchWord, m *FullMap, count *CorpusCount, docs *Corpus) []Hit {
	//setting up needed terms
	var myMap = m.getMap()
	var totalDocs = count.getCount()
	var sliceOfDocs = docs.getDoc()
	var search = nW.getWord()
	var hits []Hit

	//stem search term
	stemmed, _ := snowball.Stem(search, "english", false)
	innerMap := myMap[stemmed]

	//ranging through each document in the wanted word map
	for url, hit := range innerMap {
		//initiating the document wordCount
		var wordCount int

		//checking for the wanted document
		for _, doc := range sliceOfDocs {

			//by the url getting the word count
			if doc.getUrl() == url {
				wordCount = doc.getWordCount()
			}
		}

		//compute TF
		tf := float64(hit) / float64(wordCount)

		//compute idf
		idf := math.Log10(float64(totalDocs) / float64(len(innerMap)+1))

		//calculate TFIDF sckre
		tfidf := float64(tf) * float64(idf)

		//adding a new hit
		hits = append(hits, Hit{URL: url, Count: tfidf})
	}

	//rank the documents by sorting
	sort.Sort(ByHits(hits))
	slices.Reverse(hits)

	return hits
}
