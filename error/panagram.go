package main

import "fmt"

func Panagram() bool {
	var s string
	s = "uwqohxamknblecdtzpisgvyfjr"
	var b [123]bool
	if len(s) < 26 {
		return false
		fmt.Println("1")
	}
	for _, v := range s {
		if v > 96 && v < 123 {
			b[v] = true
		}
	}
	for i := 97; i < 123; i++ {
		if b[i] == false {
			return false
			fmt.Println("2")
		}
	}
	return true
}
