package main

import (
	"log"
	"strings"
	"unicode"

	"github.com/kljensen/snowball"

	"golang.org/x/net/html"
)

func Extract(text string, url string, idx map[string]map[string]int) (ret map[string]map[string]int, word []string, hrefs []string) {
	//parsing given text
	doc, err := html.Parse(strings.NewReader(text))
	var words []string

	//exiting if something went wrong with the Parse
	if err != nil {
		log.Fatal(err)
	}

	//code taken from the GOhtml website from specifically the parse example, edited for my uses here
	var f func(*html.Node)
	f = func(n *html.Node) {
		//iterating through each node to find the ones that matches the ones that hold the href
		if n.Type == html.ElementNode && n.Data == "a" {
			//iterating through the attributes
			for _, a := range n.Attr {
				//if href found then get the value associated
				if a.Key == "href" {
					hrefs = append(hrefs, a.Val)
					break
				}
			}
			//if a TextNode is found then range through the attributes to get the values
		} else if n.Type == html.TextNode {
			//getting the parent node
			parent := n.Parent

			//checking that the parent node fits the requirements and it isn't a style node
			if parent.Type == html.ElementNode && parent.Data != "style" {

				//iterating through each word and returning if it anything but a letter or number
				f := func(d rune) bool {
					return !unicode.IsLetter(d) && !unicode.IsNumber(d)
				}

				//appending the words into the slice using FieldsFunc to allocate the correct number of space
				words = append(words, strings.FieldsFunc(n.Data, f)...)
			}

		}
		//recursively going through each child to get to the end of the tree
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	//setting slice to be returned to the slice before it gets used for the map

	for len(words) != 0 {
		current := words[0]

		//making sure  word isnt a stop word
		if !Stop(current) {
			//stem the word
			current, _ = snowball.Stem(current, "english", false)
			word = append(word, current)

			//checking to see if the key exists
			if _, exists := idx[current]; !exists {
				idx[current] = make(map[string]int)
				idx[current][url] = 1
			} else { // if it exists just increment
				idx[current][url] += 1
			}

		}

		words = words[1:]
	}

	//returning both wanted return
	ret = idx
	return
}
