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
	var name string
	r.ParseForm()
	name = r.PostFormValue("myName") // name o'zgaruvchisiga requst body sida myName ga berilgan qiymatni o'zlashtiradi bu kode bilan faqat terminal orqali qiymat bera oldim <curl -X POST -F 'myName=Azizbek' 'http://localhost:3333/hello'>
	name = r.FormValue("myName")
	a := json.NewDecoder(r.Body).Decode(&name) //bu kode orqali postmandan qiymat olib bildim M: bodiyga yozdim  "azizbek". bu kod orqali ekrandan nima kiritilsa yozib olinadi.
	if a != nil{
		fmt.Fprint(w, a)
	}
	if name == "" {
		name = "HTTP"
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("hello %v", name)) // 32, 33, 34 qatorlar ekranga chiqarish usullari
	io.WriteString(w, fmt.Sprintf("hello %v \n", name))
	fmt.Fprintf(w, "hello %v\n", name)
}
// manimcha bu kodni veb bravzer orqali kiritish uchun front end kerak