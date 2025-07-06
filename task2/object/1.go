package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func print(s Shape) {
	fmt.Printf("面积: %v\n", s.Area())
	fmt.Printf("周长: %v\n", s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 2, Height: 4}
	fmt.Println("矩形:")
	print(rect)

	circle := Circle{Radius: 10}
	fmt.Println("圆形:")
	print(circle)
}
