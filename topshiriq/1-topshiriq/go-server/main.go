package main

import (
	"fmt"
	
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)


	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}

}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}
func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{					// post, put, 	patch requestlari hamda requestning tanasini o'qish uchun
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	n := r.FormValue("name")
	a := r.FormValue("address")
	fmt.Println(r)
	fmt.Fprintf(w, "Name = %s \n", n)
	fmt.Fprintf(w, "Address = %s", a)


}