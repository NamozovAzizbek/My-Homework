package main

type MyError struct {
	s string
	b bool
}

func (m MyError) Error() string {
	if m.b {
		return m.s
	}
	return "unknown error"
}
func GetMyError() error {
	x := MyError{"somthing wrong",true}
	return x
}
