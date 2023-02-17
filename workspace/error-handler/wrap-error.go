package main

import (
	"errors"
	"fmt"
)

type FooError struct {
	s string
	b bool
}

func (f *FooError) Error() string {
	return fmt.Sprintf("<%s, %t>", f.s, f.b)
}

type WrapError struct {
	s string
	e error
}

func (w *WrapError) Error() string {
	return fmt.Sprintf("%s // %v", w.s, w.e)
}

func (w *WrapError) Unwrap() error {
	return w.e
}
// shu funk ishlaganda yozuv chiqdi Caught FooError: <cat, true>
// ishlamaganda yozuv chiqdi        Unknown error type
var (
	errFoo = &FooError{"foo", false}
	errBar = errors.New("bar")
)

func MyWrap() {
	var errs = []error{
		errFoo,
		&FooError{"hello", true},
		errBar,
		fmt.Errorf("no wrap: %s", errFoo),
		fmt.Errorf("in fmt box: %w", errFoo),
		&WrapError{"custem wrapper", &FooError{"cat", true}},
	}
	for _, e := range errs {
		fmt.Printf("\n### %v ###\n", e)
		tryIs(e)
		tryAs(e)
	}
}

func tryIs(e error) {
	switch {
	case errors.Is(e, errFoo):
		fmt.Println("found exact errFoo in the chain")
	case errors.Is(e, errBar):
		fmt.Println("found exact errBar in the chain")
	default:
		fmt.Println("no exac match")
	}
}

func tryAs(e error){
	var foo *FooError
	switch{
	case errors.As(e, &foo):
		fmt.Printf("Caught FooError: %v", foo)
	default:
		fmt.Println("Unknown error type")
	}
}
