package main

type Gof interface {
	Sum() string
}
type Go struct{}

func (g Go) Sum() string {
	return "sdkkddkkd"
}

type MyError struct{}

func (m MyError) Error() string {
	return "boom"
}
func Get() string {
	x := Go{}
	return x.Sum()

}
