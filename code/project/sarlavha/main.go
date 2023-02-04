package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", getHello)
	http.ListenAndServe(":3333", nil)

}
func getHello(w http.ResponseWriter, r *http.Request) {
	var name string
	name = r.PostFormValue("myName")
	//_ = json.NewDecoder(r.Body).Decode(&name)
	if name == ""{
		w.Header().Set("X-Missing-Filed", "myName")//bu kodning bo'lmasligi veb saytda ko'rinmaydi, qo'shtirnoq ichiga ixtiyoriy xabar yozish mumkin shu xabarni qaytaradi
		w.WriteHeader(http.StatusBadRequest) // bu kod veb saytga 404 xatosini berarkan
		return
	}
	io.WriteString(w, name)
}
// xato kod yozilganda curl -v -X POST 'http://localhost:3333/hello'

// agar w.Header().Set("x-missing-filed", "myName") kod ishlamasa konslda
//*   Trying 127.0.0.1:3333...
// * Connected to localhost (127.0.0.1) port 3333 (#0)
// > POST /hello HTTP/1.1
// > Host: localhost:3333
// > User-Agent: curl/7.87.0
// > Accept: */*
// > 
// * Mark bundle as not supporting multiuse
// < HTTP/1.1 400 Bad Request
// < Date: Sat, 04 Feb 2023 13:25:03 GMT
// < Content-Length: 0
// < 
// * Connection #0 to host localhost left intact 

// agar w.WriteHeader(http.StatusBadRequest) kod ishlamasa konslda
// *   Trying 127.0.0.1:3333...
// * Connected to localhost (127.0.0.1) port 3333 (#0)
// > POST /hello HTTP/1.1
// > Host: localhost:3333
// > User-Agent: curl/7.87.0
// > Accept: */*
// > 
// * Mark bundle as not supporting multiuse
// < HTTP/1.1 200 OK
// < X-Missing-Filed: myName
// < Date: Sat, 04 Feb 2023 13:29:49 GMT
// < Content-Length: 0
// < 
// * Connection #0 to host localhost left intact

//agar ikkala kod ishlasa 
// *   Trying 127.0.0.1:3333...
// * Connected to localhost (127.0.0.1) port 3333 (#0)
// > POST /hello HTTP/1.1
// > Host: localhost:3333
// > User-Agent: curl/7.87.0
// > Accept: */*
// > 
// * Mark bundle as not supporting multiuse
// < HTTP/1.1 400 Bad Request
// < X-Missing-Filed: myName
// < Date: Sat, 04 Feb 2023 13:32:52 GMT
// < Content-Length: 0
// < 
// * Connection #0 to host localhost left intact