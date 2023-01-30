package main

import (
	"fmt"
	"net/http"
)
func  Main2()  {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		for i, v  := range r.Form{
			fmt.Fprint(w, i , " : ", v)
			fmt.Println(i, v)
		}
		fmt.Fprint(w,"ro'yxatdan muofaqiyatli o'tildi.")
	})
	http.ListenAndServe(":8080", mux)
}