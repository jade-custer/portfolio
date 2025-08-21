package main

import (
	"log"
	"net/url"
	"strings"

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
			for _, b := range n.Attr {
				words = append(words, b.Val)
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
	// range over urls
	//u, err := url.Parse(href) //takes a string and produces url struct
	//..... make it an absolute url including scheme and hostNamehost
	//u.String()// takes a struct and produces the string version
	var output []string

	u, err := url.Parse(host)

	if err != nil {
		log.Fatal(err)
	}

	scheme := u.Scheme
	hostName := u.Host
	for _, newHref := range href {
		newParse, err := url.Parse(newHref)
		if err != nil {
			log.Fatal(err)
		}
		hasScheme := newParse.Scheme
		if hasScheme != "https" {
			newUrl := scheme + "://" + hostName + newHref
			output = append(output, newUrl)
		} else {
			output = append(output, newHref)
		}

	}
	return output
}
