package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
mux := http.NewServeMux()

mux.HandleFunc("/", getRoot)
mux.HandleFunc("/about", about)

err := http.ListenAndServe(":8000", mux)

if errors.Is(err, http.ErrServerClosed){
	fmt.Println("Server closed !")
}else if err != nil {
	fmt.Printf("Server started ! but %v", err)
	os.Exit(1)
}
}

func getRoot(w http.ResponseWriter, r *http.Request){
	hasOne := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	fmt.Printf("got / request \n")
	io.WriteString(w, "This my website !")
}
func about(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / about request\n")
	io.WriteString(w, "Name : Azizbek")
}