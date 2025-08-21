package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestTfIdf(t *testing.T) {
	tests := []struct {
		hostUrl string
		word    string
		output  []Hit
	}{
		{
			"https://cs272-f24.github.io/top10/The%20Project%20Gutenberg%20eBook%20of%20Alice%E2%80%99s%20Adventures%20in%20Wonderland,%20by%20Lewis%20Carroll/index.html",
			"Alice",
			[]Hit{
				{"https://cs272-f24.github.io/top10/The%20Project%20Gutenberg%20eBook%20of%20Alice%E2%80%99s%20Adventures%20in%20Wonderland,%20by%20Lewis%20Carroll/index.html", -0.00013426268390930198},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap11.html", -0.000797579617075541},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap12.html", -0.0010438275688022026},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap04.html", -0.0011690236865889213},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap02.html", -0.0012291180354000474},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap10.html", -0.00124095172203164},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap03.html", -0.0012916929781164378},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap01.html", -0.0014337262945785499},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap08.html", -0.0015307816110060044},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap05.html", -0.0016003433720586252},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap06.html", -0.0016719316037091556},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap07.html", -0.0020188221173858436},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap09.html", -0.0021111240330250816},
			},
		},
		{
			"https://cs272-f24.github.io/top10/The%20Project%20Gutenberg%20eBook%20of%20Alice%E2%80%99s%20Adventures%20in%20Wonderland,%20by%20Lewis%20Carroll/index.html",
			"Caterpillar",
			[]Hit{
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap05.html", 0.014902357800056998},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap10.html", 0.0005161360049388284},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap04.html", 0.0004710253665957071},
				{"https://cs272-f24.github.io/top10/The%20Project%20Gutenberg%20eBook%20of%20Alice%E2%80%99s%20Adventures%20in%20Wonderland,%20by%20Lewis%20Carroll/index.html", 0.00024730235278356253},
			},
		},
		{
			"https://cs272-f24.github.io/top10/The%20Project%20Gutenberg%20eBook%20of%20Alice%E2%80%99s%20Adventures%20in%20Wonderland,%20by%20Lewis%20Carroll/index.html",
			"Rabbit",
			[]Hit{
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap04.html", 0.0029003558296029396},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap01.html", 0.002207845753928736},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap11.html", 0.001862400499912675},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap12.html", 0.001726495598567696},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap08.html", 0.001139363920576779},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap02.html", 0.0009035408365912977},
				{"https://cs272-f24.github.io/top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap10.html", 0.00019863288913869636},
				{"https://cs272-f24.github.io/top10/The%20Project%20Gutenberg%20eBook%20of%20Alice%E2%80%99s%20Adventures%20in%20Wonderland,%20by%20Lewis%20Carroll/index.html", 0.00019034665419250522},
			},
		},
	}

	for _, test := range tests {
		//creating idx
		idx := newMap()

		//setting up word
		nW := newWord()
		nW.changeWord(test.word)

		//setting up corpus count and hard coding it
		corp := newCount()

		// //setting up slice of documents for word count
		slice := newDocSlice()

		//getting the body of the html
		resp, err := http.Get(test.hostUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		//crawling to get the map
		m := Crawl(data, "The%20Project%20Gutenberg%20eBook%20of%20Alice%E2%80%99s%20Adventures%20in%20Wonderland,%20by%20Lewis%20Carroll/index.html", "https://cs272-f24.github.io/top10/", corp, slice)

		idx.updateMap(m)

		//checking if the expected and wanted are the same
		got := TfIdf(nW, idx, corp, slice)

		check := reflect.DeepEqual(got, test.output)

		if check != true {
			t.Errorf("TFIDF() gave %v but we wanted %v", got, test.output)
		}

	}

}
