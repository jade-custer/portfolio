package main

// creating a struct for the map
type FullMap struct {
	idx map[string]map[string]int
}

// creating a new map
func newMap() *FullMap {
	return &FullMap{make(map[string]map[string]int)}
}

// update the map
func (m *FullMap) updateMap(i map[string]map[string]int) {
	m.idx = i
}

// get the map
func (m *FullMap) getMap() map[string]map[string]int {
	return m.idx
}

// keeps track of the documents in the corpus
type Corpus struct {
	sliceOfDocs []Document
}

// creates a new slice
func newDocSlice() *Corpus {
	return &Corpus{}
}

// add a new document
func (corp *Corpus) addDoc(doc Document) {
	corp.sliceOfDocs = append(corp.sliceOfDocs, doc)
}

// get the slice of documents
func (corp *Corpus) getDoc() []Document {
	return corp.sliceOfDocs
}

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

// getting the number of documents in the corpus
type CorpusCount struct {
	count int
}

// creating a new count
func newCount() *CorpusCount {
	return &CorpusCount{}

}

// updating the count
func (c *CorpusCount) changeCount(count int) {
	c.count = count
}

// get the count
func (c *CorpusCount) getCount() int {
	return c.count
}

// searchword struct
type SearchWord struct {
	word string
}

// creating a new search word
func newWord() *SearchWord {
	return &SearchWord{}

}

// changing the word
func (p *SearchWord) changeWord(newWord string) {
	p.word = newWord
}

// getting the word
func (p *SearchWord) getWord() string {
	return p.word
}
