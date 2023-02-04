package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", getHello)

	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed %v", err)
		return
	} else if err != nil {
		fmt.Printf("Server stared but : %s", err)
	}
}
func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("myName") // name o'zgaruvchisiga requst body sida myName ga berilgan qiymatni o'zlashtiradi
	json.NewEncoder(w).Encode(fmt.Sprintf("hello %v", name)) // 32, 33, 34 qatorlar ekranga chiqarish usullari
	io.WriteString(w, fmt.Sprintf("hello %v \n", name))
	fmt.Fprintf(w, "hello %v\n", name)
}
// bu kodga faqat terminalda murojat qilib bildim postman va veb bravzerdan qiymani kiritib bilmadim  <curl -X POST -F 'myName=Azizbek' 'http://localhost:3333/hello'>