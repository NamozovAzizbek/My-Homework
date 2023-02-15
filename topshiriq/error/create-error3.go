package main

import "fmt"

type IntError int

func (i IntError) Error() string{
	return fmt.Sprintf("Error #%d", int(i))
}

func gerIntError() error{
	return IntError(1)
}