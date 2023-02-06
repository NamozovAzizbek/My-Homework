package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.Write([]byte(`{"hello": "world"}`))
	//w.Write([]byte(`<?xml version="1.0" encoding="utf-
    //8"?><Message>Hello World</Message>`))
}

func Main1() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}
