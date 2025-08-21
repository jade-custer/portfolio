package main

type Map interface {
	CreateMap() 
	AddNewDoc() 
	URLExists(url string) bool
	IncrementWordCount(word string) 
	AddToMap(word string, url string, hit int)
	QueryInnerMap(word string) []Hit
	QueryWordCount(url string) int
	QueryDocCount() int
}