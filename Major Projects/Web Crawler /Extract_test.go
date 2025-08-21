package main

import (
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {
	//creating test struct
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

		{
			`<html>
   			 <head>
        	<title>The </title>
    		</head>
    		<body>
        	<p>Hello World!</p>
        	<p>Welcome to <a href="https://cs272-f24.github.io/">CS272</a>!</p>
    		</body>
			</html>`,

			[]string{"Hello", "World!", "Welcome", "to", "CS272"},
			[]string{"https://cs272-f24.github.io/"},
		},
	}

	//looping through each test case
	for _, test := range tests {
		//needed to initiate an empty map to get rid of errors
		idx := make(map[string]map[string]int)
		_, words, hrefs := Extract(test.doc, "", idx)

		//checking if the given and wanted are the same
		gotWords := reflect.DeepEqual(words, test.wantedWords)
		gotHrefs := reflect.DeepEqual(hrefs, test.WantedHrefs)

		if !gotWords && !gotHrefs {
			t.Errorf("Extract()) gave %v and %v when %v and %v were wanted", gotWords, gotHrefs, test.wantedWords, test.WantedHrefs)
		}
	}

}
