package main

import "fmt"

type IntError int

func (i IntError) Error() string {
	return fmt.Sprintf("Error #%d", i)
}
func GetIntError() error{
	 var x IntError
	 x = 1
	 fmt.Printf("%T, %v, %T, %s\n", x, x, x.Error(), x.Error())
	 return x
}
