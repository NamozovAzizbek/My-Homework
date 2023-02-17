package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)
type RequestError struct {
	StatusCode int
	Err        error
}

func (r RequestError) Error() string {
	return r.Err.Error()
}

func (r *RequestError) Temporary() bool {
	return r.StatusCode == http.StatusServiceUnavailable // 503
}

func doRequest() error {
	return RequestError{
		StatusCode: 503,
		Err:        errors.New("unaviable"),
	}
}
func CheckError(){
	err := doRequest()
	if err != nil{
		fmt.Println(err)
		if req, ok := err.(RequestError); ok{
			if req.Temporary(){
				fmt.Println("This request can be tried again")
			}else {
				fmt.Println("This request cannot be tried again")
			}
		}
		os.Exit(1)
	}
	fmt.Println("succes !")
}