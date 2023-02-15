package main

type MyError struct{
	s string
	b bool
}

func (e MyError) Error() string{
	if e.b{
		return e.s
	}
	return "unknown error"
}

func getMyError() error{
	return MyError{"something", true}
}