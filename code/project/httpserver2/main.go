// bu kodda serverga murojat qilinganda querilardagi parametrlarning qiymatlari o'zgaruvchilarga o'zlashtiriladi.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"

	"net/http"
	"os"
)

var serverAddr string

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/about", about)

	ctx := context.Background()

	server := &http.Server{ // bu qatorda server 3333 portda bo'lishi elon qilindi
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context { // bu qatorda elon qilingan fulksiya orqali getRootda uning addressi olindi bu funksiyasiz ham serverni elon qilish munkin lekin uning qaysi portda xizmat qilayotgani bilib bo'lmaydi.
			ctx = context.WithValue(ctx, serverAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe() // bu qatorda yuqorida elon qilingan serverimiz yurg'azildi.

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed !")
	} else if err != nil {
		fmt.Printf("Server started ! but %v", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hasOne := r.URL.Query().Has("first") // bu kod birinchi parametrga qiymat berilganmi yoqligin qaytaradi (true || false)
	fir := r.URL.Query().Get("first")    // firdegan o'zgaruvchiga queryda first=qiymat berilgandagi qiymatni o'zlashtiradi.
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	fmt.Printf("%v got / request first(%t) %s, second(%t) %s \n",
		ctx.Value(serverAddr), //buserver xizmat ko'rsatayotgan port manzilin olib beradi.
		hasOne, fir,
		hasSecond, second)
	io.WriteString(w, "This my website !")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / about request\n")
	io.WriteString(w, "Name : Azizbek")
}

// funksiyaga brazerdan murojat qilinganda parametrlardan keladigan qiymatlarni o'zlashtirmaydi. Parametrlarga berilgan qiymatni olishni terminalda <curl 'http://localhost:3333?first=1&second=2'> kodni yozish orqali olish mumkin

// bravzerda parametrlarga berilgan qiymatlarni olish uchun <http://localhost:3333/?first=1&second=2> kod yoziladi. farqi bunda port raqimidan keyin </> slesh belgisin bor.
