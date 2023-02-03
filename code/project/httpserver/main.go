// bu kodda vebsayt ikkita portda xizmat qiladi 8000 va 3333 portlar
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

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s : got / request \n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This my  website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s : got /hello request \n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello HTTP!\n")
}

func main() {
	//		http.HandleFunc("/", getRoot)
	//		http.HandleFunc("/hello", getHello)
	//		err := http.ListenAndServe(":8000", nil)
	//		if errors.Is(err, http.ErrServerClosed){
	//			fmt.Print("server closed !")
	//	}else if err != nil {
	//
	//		fmt.Printf("server starting ! but %v\n", err)
	//		os.Exit(1)
	//	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":8000",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	serverTwo := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed !")
		} else if err != nil {
			fmt.Printf("Server starting ! but %v", err)
			os.Exit(1)
		}
		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed !")
		} else if err != nil {
			fmt.Printf("Server starting ! but %v", err)
			os.Exit(1)
		}
		cancelCtx()
	}()
	<-ctx.Done()
}
