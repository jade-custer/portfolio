package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

func Extract(text string) (words []string, hrefs []string) {
	//parsing given text
	doc, err := html.Parse(strings.NewReader(text))

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

	//returning both wanted []strings
	return
}

func Clean(host string, href []string) []string {
	//initializing variables
	var output []string

	//parsing the host and checking for errors
	u, err := url.Parse(host)

	if err != nil {
		log.Fatal(err)
	}

	//getting scheme and host from url
	scheme := u.Scheme
	hostName := u.Host

	//ranging through hrefs and adding completed urls
	for _, newHref := range href {
		//getting newHref and checking for errors
		newParse, err := url.Parse(newHref)
		if err != nil {
			log.Fatal(err)
		}

		//checking the scheme
		hasScheme := newParse.Scheme

		//if not a new link then add host, scheme and new href
		if hasScheme != "https" {
			newUrl := scheme + "://" + hostName + newHref
			output = append(output, newUrl)
			//if a new url as is
		} else {
			output = append(output, newHref)
		}

	}

	return output
}

func Download(url string) ([]byte, error) {
	//initalizing variables
	var err error
	//using http.Get to get the body test and return it
	if rsp, err := http.Get(url); err == nil {
		defer rsp.Body.Close()
		if bts, err := io.ReadAll(rsp.Body); err == nil {
			return bts, nil
		}
	}
	return []byte(""), err

}

func Crawl(seedUrl string) (finalRangeLength int) {
	//initializing  queue and getting elements from first page
	var queue []string
	var boolean bool
	var finalRange []string
	queue = append(queue, seedUrl)
	tempUrl, _ := url.Parse(seedUrl)
	hostName := tempUrl.Host
	finalRange = append(finalRange, seedUrl)

	//extracting urls and hrefs
	for len(queue) != 0 {
		current := queue[0]

		//downloading seed url
		downloaded, err := Download(current)

		//checking for errors
		if err != nil {
			log.Fatal(err)
		}

		//performing a crawl
		_, newHrefs := Extract(string(downloaded))
		cleanedUrls := Clean(current, newHrefs)
		finalRange = append(finalRange, cleanedUrls...)

		//checking the new urls for repeats and not host links
		for len(cleanedUrls) != 0 {
			cur := cleanedUrls[0]
			u, _ := url.Parse(cur)

			//if an outside link don't crawl
			if u.Host != hostName {
				fmt.Printf("not part of seed ")

			} else {
				//checking if it has been crawled or will be crawled
				for _, exists := range finalRange {
					if exists == cur {
						boolean = true
					}
				}

				//if not crawled/about to be crawled add it to the queue
				if !boolean {
					queue = append(queue, cur)
				}
			}
			//creating a new slice without the first element
			cleanedUrls = cleanedUrls[1:]
		}
		//popping off the first element
		queue = queue[1:]
	}

	return len(finalRange)
}
