package main

import (
	"fmt"
)

type Shape interface {
	Area() float32
}

type Square struct {
	Size float32
}

type Rectangle struct {
	Width, Height float32
}

type Triangle struct {
	Width, Height float32
}

func (s Square) Area() float32 {
	return s.Size * s.Size
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Width * t.Height
}

func TotalArea(shapes []Shape) float32 {
	var total float32
	for _, shape := range shapes{
		total += shape.Area()
	}
	return total
}

func ExampleTotalArea() {
	fmt.Println(TotalArea([]Shape{
		Square{3},
		Rectangle{4, 5},
		Triangle{6, 7},
	}))
}

func main() {
	ExampleTotalArea()
}
