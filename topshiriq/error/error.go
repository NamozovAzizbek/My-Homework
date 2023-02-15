package main

type error interface {
	Error() string
}
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error { // New text formatida error qaytaradi
	return &errorString{text}
}
func Main() {
	var a errorString
	a.s = "ukhjhkhjj"
	error.Error(&errorString{})
}
