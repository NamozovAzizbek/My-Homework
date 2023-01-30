package main

import (
	"fmt"
	"net/http"
)

func Value(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", getParams)

	http.ListenAndServe(":8080", mux)

}
func getParams(w http.ResponseWriter, r *http.Request){
	a := r.FormValue("a")
	fmt.Fprint(w, a)
	w.Write([]byte(a))
	//agarda brovzerga yozsang http://localhost:8080/?a= malumot  -> malumitni ekranga chiqaradi
}