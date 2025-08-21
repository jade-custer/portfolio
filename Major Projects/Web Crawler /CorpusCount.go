package main

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
