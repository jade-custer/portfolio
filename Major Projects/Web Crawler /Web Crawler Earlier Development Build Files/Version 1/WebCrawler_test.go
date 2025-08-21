package main

import (
	"net/http"
	"net/http/httptest"
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
	}

	//looping through each test case
	for _, test := range tests {
		words, hrefs := Extract(test.doc)

		//checking if the given and wanted are the same
		gotWords := reflect.DeepEqual(words, test.wantedWords)
		gotHrefs := reflect.DeepEqual(hrefs, test.WantedHrefs)

		if gotWords != true && gotHrefs != true {
			t.Errorf("Extract()) gave %v and %v when %v and %v were wanted", gotWords, gotHrefs, test.wantedWords, test.WantedHrefs)
		}
	}

}

func TestCleanHref(t *testing.T) {
	//initiializing test struct
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

func TestDownload(t *testing.T) {
	tests := []struct {
		expected string
		want     string
	}{
		{
			expected: `<html> <body> Hello CS 272 </body> </html>`,
			want:     `<html> <body> Hello CS 272 </body> </html>`,
		},

		{
			expected: `<html><head></head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
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
			</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>`,

			want: `<html><head></head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
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
			</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>`,
		},

		{
			expected: `<html><head></head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
			For a simple example, see <a href="/tests/project01/simple.html">simple.html</a>
			</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>`,

			want: `<html><head></head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
			For a simple example, see <a href="/tests/project01/simple.html">simple.html</a>
			</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>`,
		},

		{
			expected: `<html><head></head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
			Hello CS 272, there are no links here.
			</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>`,

			want: `<html><head></head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
			Hello CS 272, there are no links here.
			</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>`,
		},

		{
			expected: `<html><head>
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

			want: `<html><head>
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
		},
	}

	for _, test := range tests {
		//creating mock server
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(test.expected))
		}))

		got, _ := Download(svr.URL)
		check := reflect.DeepEqual(string(got), test.want)

		if check != true {
			t.Errorf("Wanted %v and got %v instead", test.want, got)
		}

		defer svr.Close()
	}
}

func TestCrawl(t *testing.T) {
	//initializing test struct
	tests := []struct {
		url    string
		length int
	}{
		{
			`<body> Hello CS 272, there are no links here. </body>`,

			1,
		},

		{
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

			4,
		},
		{
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

			3,
		},
		{
			`<html><head></head><body data-scholarcy-content-script-executed="1" data-new-gr-c-s-check-loaded="14.1194.0" data-gr-ext-installed="">
			For a simple example, see <a href="/tests/project01/simple.html">simple.html</a>
			</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>`,

			2,
		},
	}

	for _, test := range tests {
		//creating mock server
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(test.url))
		}))
		len := Crawl(svr.URL)

		if len != test.length {
			t.Errorf("Wanted %v  but got  %v instead", test.length, len)
		}

		defer svr.Close()
	}

}
