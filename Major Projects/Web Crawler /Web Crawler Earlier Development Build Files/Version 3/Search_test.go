package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		url        string
		word       string
		wantedFreq int
	}{
		{
			`<p class="drama">
			CHORUS.<br/>
			Two households, both alike in dignity,<br/>
			In fair Verona, where we lay our scene,<br/>
			From ancient grudge break to new mutiny,<br/>
			Where civil blood makes civil hands unclean.<br/>
			From forth the fatal loins of these two foes<br/>
			A pair of star-cross’d lovers take their life;<br/>
			Whose misadventur’d piteous overthrows<br/>
			Doth with their death bury their parents’ strife.<br/>
			The fearful passage of their death-mark’d love,<br/>
			And the continuance of their parents’ rage,<br/>
			Which, but their children’s end, nought could remove,<br/>
			Is now the two hours’ traffic of our stage;<br/>
			The which, if you with patient ears attend,<br/>
			What here shall miss, our toil shall strive to mend.
			</p>`,
			"fear",
			1,
		},
		{
			`<p class="drama">
			CHORUS.<br/>
			Two households, both alike in dignity,<br/>
			In fair Verona, where we lay our scene,<br/>
			From ancient grudge break to new mutiny,<br/>
			Where civil blood makes civil hands unclean.<br/>
			From forth the fatal loins of these two foes<br/>
			A pair of star-cross’d lovers take their life;<br/>
			Whose misadventur’d piteous overthrows<br/>
			Doth with their death bury their parents’ strife.<br/>
			The fearful passage of their death-mark’d love,<br/>
			And the continuance of their parents’ rage,<br/>
			Which, but their children’s end, nought could remove,<br/>
			Is now the two hours’ traffic of our stage;<br/>
			The which, if you with patient ears attend,<br/>
			What here shall miss, our toil shall strive to mend.
			</p>`,
			"blood",
			1,
		},
		{
			`<p class="drama">
			CHORUS.<br/>
			Two households, both alike in dignity,<br/>
			In fair Verona, where we lay our scene,<br/>
			From ancient grudge break to new mutiny,<br/>
			Where civil blood makes civil hands unclean.<br/>
			From forth the fatal loins of these two foes<br/>
			A pair of star-cross’d lovers take their life;<br/>
			Whose misadventur’d piteous overthrows<br/>
			Doth with their death bury their parents’ strife.<br/>
			The fearful passage of their death-mark’d love,<br/>
			And the continuance of their parents’ rage,<br/>
			Which, but their children’s end, nought could remove,<br/>
			Is now the two hours’ traffic of our stage;<br/>
			The which, if you with patient ears attend,<br/>
			What here shall miss, our toil shall strive to mend.
			</p>`,
			"attend",
			1,
		},
		{
			`<p class="drama">
			CHORUS.<br/>
			Two households, both alike in dignity,<br/>
			In fair Verona, where we lay our scene,<br/>
			From ancient grudge break to new mutiny,<br/>
			Where civil blood makes civil hands unclean.<br/>
			From forth the fatal loins of these two foes<br/>
			A pair of star-cross’d lovers take their life;<br/>
			Whose misadventur’d piteous overthrows<br/>
			Doth with their death bury their parents’ strife.<br/>
			The fearful passage of their death-mark’d love,<br/>
			And the continuance of their parents’ rage,<br/>
			Which, but their children’s end, nought could remove,<br/>
			Is now the two hours’ traffic of our stage;<br/>
			The which, if you with patient ears attend,<br/>
			What here shall miss, our toil shall strive to mend.
			</p>`,
			"ancient",
			1,
		},
	}

	for _, test := range tests {
		//reading url
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(test.url))
		}))
		defer server.Close()

		//reading the server for the html text
		resp, err := http.Get(server.URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		//checking if the expected and wanted are the same
		got := Search(data, test.word, server.URL)

		check := reflect.DeepEqual(test.wantedFreq, got)
		if check != true {
			t.Errorf("Wanted %v  but got  %v instead", test.wantedFreq, got)
		}

	}

}
