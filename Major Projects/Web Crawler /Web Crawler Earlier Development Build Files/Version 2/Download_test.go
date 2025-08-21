package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

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

		//chekcing if expected matches wanted
		got, _ := Download(svr.URL)
		check := reflect.DeepEqual(string(got), test.want)

		if check != true {
			t.Errorf("Wanted %v and got %v instead", test.want, got)
		}

		defer svr.Close()
	}
}
