/*package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
func Value2(){
	r := mux.NewRouter()
	r.HandleFunc("/", han).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
func han(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	a := r.FormValue("one")
	b :=r.FormValue("two")
	log.Println(a, b)
	fmt.Fprint(w, "Gorilla !\n")
}*/

package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
//   log.Println("one = " + r.FormValue("one"))
//   log.Println("two = " + r.FormValue("two"))
o := r.FormValue("one")
t := r.FormValue("two")
  fmt.Fprintf(w, "%v Gorilla! %v \n", o, t)
}

func Value2() {
  r := mux.NewRouter()
  r.HandleFunc("/form", YourHandler).Methods("POST")
  log.Fatal(http.ListenAndServe(":8000", r))
}
