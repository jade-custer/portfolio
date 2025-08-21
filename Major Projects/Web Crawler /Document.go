package main

// creating a new document
type Document struct {
	url       string
	wordCount int
}

// creating a new document
func newDoc(url string, wC int) *Document {
	return &Document{url, wC}

}

// getting just the url from the document
func (d *Document) getUrl() string {
	return d.url
}

// getting the word count from the document
func (d *Document) getWordCount() int {
	return d.wordCount
}
