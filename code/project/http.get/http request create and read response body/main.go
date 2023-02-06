package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var serverPort = 3333

func main() {

	go func() {// serverni boshqa tritda ishga tushirishning sababi agar server bilan main bir tritda ishlaydigan bo'lsa mainda request jo'natish uchun server ishlashi kerak demak birinchi bo'lib server ni yurg'azish kerak server yurg'azganimdan keyin server ishga tushirilgan qatordan pastki qatorni main funksiyasi o'qimaydi request ham jo'ntilmaydi.
		mux := http.NewServeMux()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s\n", r.Method)
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
	time.Sleep(100 * time.Millisecond)// bu kod yuqoridagi funksiyadi http serverni boshqa tritda ish ga tushurish uchun main funksiya 100 millisecond ushlab turadi va shu muddatda server ishga tushadi va pastki qatordagi kodlar ishlashni boshlaydi. agar time.sleep qilinmasa main funksiysi birinchi bajarilib http serverni ishlatuvchi funksiya bajarilmay qolishi mumkin
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	req, err := http.NewRequest("GET", requestURL, nil)// "GET" or http.MethodGet.Bu kod orqali get methodida ishlaydigan yangi request hosil qilindi undan tashqari bu kod http.get o'xshab birdaniga request jo'natmaydi bu vazifa yani request jo'natishni http.DefaultClient orqali amalga oshiriladi
	if err != nil {
		fmt.Printf("clinet: could not creat request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)// http.DefaultClient http.get bilan bir xilda.
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	// respons ekranga bostrildi.
	readBody, err := ioutil.ReadAll(res.Body)// ioutil.ReadAll responsni o'qish uchun ishlatildi
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: respons body: %s\n", readBody)
}
