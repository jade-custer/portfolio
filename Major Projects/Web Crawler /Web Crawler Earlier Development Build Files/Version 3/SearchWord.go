package main

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
