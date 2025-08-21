package main

type InvIdx struct {
	idx map[string]map[string]int
	documents map[string]int
}

func (i *InvIdx) CreateMap(){
	i.idx = make(map[string]map[string]int)
	i.documents = make(map[string]int)
}

	// AddNewDoc() 
	// URLExists(url string) bool
	// IncrementWordCount(word string) 
	// AddToMap(word string, url string, hit int)
	// QueryInnerMap(word string) []Hit
	// QueryWordCount(url string) int
	// QueryDocCount() int
