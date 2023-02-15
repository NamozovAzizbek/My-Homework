// errorni yaratishning ikki usuli mavjud
package main

import (
	"errors"
	"fmt"
	"math"
)

func main1() {
	res, err := Sqrt1(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println(math.Sqrt(-1))// bu hollatda sqrt berilgan argument manfiylig haqida xabar berilmaydi.
	
}

// error yaratishning 1-usuli
// import "errors"
func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("math: sqrt kritilgan son manfiy")
	}
	return math.Sqrt(num), nil
}

// error yaratishning 2-usuli
// import "fmt"
func Sqrt1(num float64)(float64, error){
	if num < 0 {
		return 0, fmt.Errorf("math: sqrt kritilgan son manfiy %g", num) // argument qiymati ham qaytariladi
	}
	return math.Sqrt(num), nil
}

