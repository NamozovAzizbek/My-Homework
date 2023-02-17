package main

import "fmt"

func Generik() {
	intsToAdd := []int{10, 11, 45, 3}
	floatToAdd := []float64{2.3, 3.5, 5.9}
	fmt.Println("int sum: ", sumNambers(intsToAdd))
	fmt.Println("float sum: ", sumNambers(floatToAdd))
}

// func sumFloats(nums []float64) float64 {
// 	var res float64
// 	for _, v := range nums {
// 		res += v
// 	}
// 	return res
// }

// func sumInts(nums []int) int {
// 	var res int
// 	for _, v := range nums {
// 		res += v
// 	}
// 	return res
// }
// kamentdagi ikkita funksiyani shunaqa qib yozsa bo'larkan
type Number interface{
	int | int64 | float32 | float64
}
func sumNambers[n Number](num []n) n{
	var res n
	for _, v := range num{
		res += v
	}
	return res
}