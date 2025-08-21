package main

import (
	"reflect"
	"testing"
)

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
		//checking the given output vs the expected
		got := Clean(test.hostName, test.hrefs)
		compare := reflect.DeepEqual(got, test.want)
		if compare != true {
			t.Errorf("Clean() gave %v but we wanted %v", got, test.want)
		}

	}
}
