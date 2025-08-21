package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawl(t *testing.T) {
	//initializing test struct
	tests := []struct {
		word string
		body string
		hit  int
	}{
		{
			"272",
			`<body> Hello CS 272, there are no links here. </body>`,
			1,
		},
		{
			"href",
			`<body>
		  	<ul>
		    <li>
		    <a href="/tests/project01/simple.html">simple.html</a>
		    </li>
		    <li>
		    <a href="/tests/project01/href.html">href.html</a>
		    </li>
		    <li>
		    <a href="/tests/project01/style.html">style.html</a>
		  	</li></ul>
			</body>`,

			1,
		},
		{
			"blue",
			`<html><head>
			<title>Style</title>
			<style>
			a.blue {
			color: blue;
			}
			a.red {
			color: red;
			}
			</style>
			</head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
			<p>
			Here is a blue link to <a class="blue" href="/tests/project01/href.html">href.html</a>
			</p>
			<p>
			And a red link to <a class="red" href="/tests/project01/simple.html">simple.html</a>`,

			1,
		},
		{
			"simpl",
			`<html><head></head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
			For a simple example, see <a href="/tests/project01/simple.html">simple.html</a>
			</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>`,

			2,
		},
	}

	for _, test := range tests {
		//initiating needed structs to let Crawl run
		corp := newCount()
		slice := newDocSlice()

		//creating a mock server to run this test
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(test.body))
		}))
		defer server.Close()

		//getting the body of the html
		resp, err := http.Get(server.URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		//crawling to get the map
		m := Crawl(data, "", server.URL, corp, slice)

		//checking if the expected output is the same as the wanted
		if m[test.word][server.URL] != test.hit {
			t.Errorf("Wanted %v  but got  %v instead", test.hit, m[test.word][server.URL])
		}
	}

}
