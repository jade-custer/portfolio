package main

import (
	"io"
	"net/http"
)

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
