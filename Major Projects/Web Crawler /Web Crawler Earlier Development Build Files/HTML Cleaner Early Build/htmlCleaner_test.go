package main

import (
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		doc         string
		wantedWords []string
		WantedHrefs []string
	}{
		{
			`<body>
  			<p>Some text here</p>
  			<a href="http://example.com">Example</a>
			</body>`,

			[]string{"Some", "text", "here", "Example"},
			[]string{"http://example.com"},
		},

		{
			`<html>
   			 <head>
        	<title>CS272 | Welcome</title>
    		</head>
    		<body>
        	<p>Hello World!</p>
        	<p>Welcome to <a href="https://cs272-f24.github.io/">CS272</a>!</p>
    		</body>
			</html>`,

			[]string{"CS272", "|", "Welcome", "Hello", "World!", "Welcome", "to", "CS272"},
			[]string{"https://cs272-f24.github.io/"},
		},
	}

	//looping through each test case
	for _, test := range tests {
		words, hrefs := Extract(test.doc)
		gotWords := reflect.DeepEqual(words, test.wantedWords)
		gotHrefs := reflect.DeepEqual(hrefs, test.WantedHrefs)

		if gotWords != true && gotHrefs != true {
			t.Errorf("Extract()) gave %v and %v when %v and %v were wanted", gotWords, gotHrefs, test.wantedWords, test.WantedHrefs)
		}
	}

}

func TestCleanHref(t *testing.T) {
	//clean test
	tests := []struct {
		hostName string   //hostName to start at, and pre-pend to partial
		hrefs    []string // input hrefs, could be absolute or partial
		want     []string //expected output, absolute URLS
	}{
		{
			"https://CS272.com",
			[]string{"/", "/documents/"},
			[]string{"https://CS272.com/", "https://CS272.com/documents/"},
		},

		{
			"https://cs272-f24.github.io/",
			[]string{"/", "/help/", "/syllabus/", "https://gobyexample.com/"},
			[]string{"https://cs272-f24.github.io/", "https://cs272-f24.github.io/help/", "https://cs272-f24.github.io/syllabus/", "https://gobyexample.com/"},
		},

		{
			"",
			nil,
			nil,
		},
	}

	for _, test := range tests {
		got := Clean(test.hostName, test.hrefs)
		compare := reflect.DeepEqual(got, test.want)
		if compare != true {
			t.Errorf("Clean() gave %v but we wanted %v", got, test.want)
		}

	}
}
