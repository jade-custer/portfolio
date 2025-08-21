package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Crawl(seedUrl []byte, org string, orgUrl string, corp *CorpusCount, slice *Corpus) (idx map[string]map[string]int) {
	//initializing  queue and other slices
	var queue []string
	queue = append(queue, org)
	var hrefs []string
	var newSlice []string
	idx = make(map[string]map[string]int)

	//looping through all links
	for len(queue) != 0 {
		//adding another doc to the count
		corp.changeCount(corp.count + 1)

		//getting the url
		current := queue[0]
		//log.Println(current)

		if strings.HasPrefix(current, ":///top10/") {
			current = strings.TrimPrefix(current, ":///top10/")
		}

		resp, err := http.Get(orgUrl + current)

		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		//performing a crawl
		idx, newSlice, hrefs = Extract(string(data), orgUrl+current, idx)
		d := newDoc(orgUrl+current, len(newSlice))
		slice.addDoc(*d)
		cleaned := Clean("", hrefs)

		//checking to see if these need to be added into queue
		for len(cleaned) != 0 {
			url := cleaned[0]
			var firstKey string

			if strings.HasPrefix(url, ":///top10/") {

				//getting an arbitrary element
				for key := range idx {
					firstKey = key
					break // Exit the loop after the first element
				}

				if subMap, exists := idx[firstKey]; exists {
					if _, exists := subMap[url]; !exists {
						queue = append(queue, url)
					}

				}
			}

			cleaned = cleaned[1:]
		}

		//popping off the first element
		queue = queue[1:]
	}

	return
}
