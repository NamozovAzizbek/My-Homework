package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var serverPort = 3333

func main() {
	go func() { // serverni boshqa tritda ishga tushirishning sababi agar server bilan main bir tritda ishlaydigan bo'lsa mainda request jo'natish uchun server ishlashi kerak demak birinchi bo'lib server ni yurg'azish kerak server yurg'azganimdan keyin server ishga tushirilgan qatordan pastki qatorni main funksiyasi o'qimaydi request ham jo'ntilmaydi.
		mux := http.NewServeMux()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s %v\n", r.Method, r.URL.Path)
			fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
			fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
			fmt.Printf("server: headers:\n")
			for headerName, headerValue := range r.Header {
				fmt.Printf("%s = %s\n", headerName, headerValue)
			}
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("server: could not read request body: %s\n", err)
			}
			fmt.Printf("server: request body: %s\n", reqBody)
			fmt.Fprintf(w, `{"message":"hello!"}`)
		})
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server %s\n ", err)
			}
		}
	}()
	time.Sleep(100 * time.Millisecond) // bu kod yuqoridagi funksiyadi http serverni boshqa tritda ish ga tushurish uchun main funksiya 100 millisecond ushlab turadi va shu muddatda server ishga tushadi va pastki qatordagi kodlar ishlashni boshlaydi. agar time.sleep qilinmasa main funksiysi birinchi bajarilib http serverni ishlatuvchi funksiya bajarilmay qolishi mumkin
	jsonBody := []byte(`{"client_message": "hello, server!"}`)
	bodyReader := bytes.NewReader(jsonBody)
	requestURL := fmt.Sprintf("http://localhost:%d?id=123", serverPort)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("clinet: error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	readResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read respons %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", readResponse)
}
