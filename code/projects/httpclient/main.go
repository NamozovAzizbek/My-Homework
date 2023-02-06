package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

const serverPort = 3333

func main() {
	 // server()
	go func() { // serverni boshqa tritda ishga tushirishning sababi agar server bilan main bir tritda ishlaydigan bo'lsa mainda request jo'natish uchun server ishlashi kerak demak birinchi bo'lib server ni yurg'azish kerak server yurg'azganimdan keyin server ishga tushirilgan qatordan pastki qatorni main funksiyasi o'qimaydi request ham jo'ntilmaydi.
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %v /\n", r.Method)
		})
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
		fmt.Println("server runnig")
	}()
	time.Sleep(100 * time.Millisecond)// bu kod yuqoridagi funksiyadi http serverni boshqa tritda ish ga tushurish uchun main funksiya 100 millisecond ushlab turadi va shu muddatda server ishga tushadi va pastki qatordagi kodlar ishlashni boshlaydi. agar time.sleep qilinmasa main funksiysi birinchi bajarilib http serverni ishlatuvchi funksiya bajarilmay qolishi mumkin
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: got response ! \n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	
}
// func server() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("server: %v /\n", r.Method)
// 	})
// 	server := &http.Server{
// 		Addr:    fmt.Sprintf(":%d", serverPort),
// 		Handler: mux,
// 	}
// 	if err := server.ListenAndServe(); err != nil {
// 		if !errors.Is(err, http.ErrServerClosed) {
// 			fmt.Printf("error running http server: %s\n", err)
// 		}
// 	}
// 	// fmt.Println("server runnig")
// }
