package main

import (
	"github.com/kljensen/snowball"
)

func Search(url []byte, wantedWord string, orgUrl string) int {
	//initiating needed structs for the crawl
	corp := newCount()
	slice := newDocSlice()

	//crawl data structure
	idx := Crawl(url, "", orgUrl, corp, slice)

	//stem search term
	stemmed, _ := snowball.Stem(wantedWord, "english", false)

	//returning frequency
	return idx[stemmed][orgUrl]

}
