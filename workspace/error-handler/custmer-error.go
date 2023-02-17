package main

import (
	"errors"
	"fmt"
)

type DivisionError struct{
	Number1, Number2 int
	Msg string
}

func (d *DivisionError) Error() string{
	return d.Msg
}

func DivideCustmer(a, b int) (int, error){
	if b == 0{
		return 0, &DivisionError{
			Msg: fmt.Sprintf("cannot divide '%d' by zero", a),
			Number1: a,
			Number2: b,
		}
	}
	return a/b, nil
}
func Check(){
	a, b := 10, 0
	res, err := DivideCustmer(a, b)
	if err != nil{
		var divErr *DivisionError
		switch{
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is not mathematically valid: %s\n",
			divErr.Number1, divErr.Number2, divErr.Error())
		default:
            fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}
	fmt.Printf("%d / %d = %d\n", a, b, res)
}