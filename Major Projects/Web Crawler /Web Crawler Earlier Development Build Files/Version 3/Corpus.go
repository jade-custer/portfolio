package main

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
