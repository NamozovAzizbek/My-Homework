package main

import (
	"fmt"
	"strings"
)

func Trim() {
	fmt.Println("Please enter your name.")
	var name, fulname string
	fmt.Scanln(&name)
	fmt.Scanf("%s",&fulname)
	name = strings.TrimSpace(name)
	fmt.Printf("Hi, %s, %s! I'm Go!\n", name, fulname)
}
