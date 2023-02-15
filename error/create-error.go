package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func Create() {
fmt.Println("math : ",math.Sqrt(-1))
_, err := Sqrt1(-1)
fmt.Println(err)
os.Exit(1)
}
// import "error"
func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("math: sqrt kritilgan son manfiy\n")
	}
	return math.Sqrt(num), nil
}
// import "fmt"
func Sqrt1(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("math: sqrt kritilgan son manfiy %g\n", num) // argument qiymati ham qaytariladi
	}
	return math.Sqrt(num), nil
}
