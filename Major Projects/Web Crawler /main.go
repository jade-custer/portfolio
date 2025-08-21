package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	//setting up instances
	nW := newWord()
	idx := newMap()
	corp := newCount()
	slice := newDocSlice()

	//setting up web server to run concurrent
	go Serve(nW, idx, corp, slice)

	//setting up the needed body
	resp, err := http.Get("http://localhost:8080/top10/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	//setting up crawl to run concurrently
	go func() {
		idx.updateMap((Crawl(data, "index.html", "http://localhost:8080/top10/", corp, slice)))
	}()

	//forcing main to not close
	for {
		time.Sleep(1000 * time.Millisecond)
	}
}
