package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimetr() float64
}

type Rectangle struct {
	Lenght, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Lenght * r.Width
}
func (r Rectangle) Perimetr() float64 {
	return 2 * (r.Lenght + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimetr() float64 {
	return 2 * math.Pi * c.Radius
}

func CalculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, sh := range shapes{
		totalArea += sh.Area()
	}
	return totalArea
}

func Run() Shape{
	return Circle{2}
}
func Interface() {

	var s = Circle{5.0}
	fmt.Println(s.Area(), s.Perimetr())

	var t Shape = Rectangle{2.0, 1.0}
	fmt.Println(t.Area(), t.Perimetr())

	var k = Rectangle{2, 2}
	fmt.Println(k.Area(), k.Perimetr())

	totalArea := CalculateTotalArea(Circle{3}, Rectangle{3,4}, Circle{5})
	fmt.Println(totalArea, CalculateTotalArea(s,t, k))
}
