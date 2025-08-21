package main

import (
	"log"
	"net/url"
)

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
