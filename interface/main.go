package main

import (
	"fmt"
	"math"
)

// Defined personal type square
type square struct {
	side float64
}

// method to calculate area for type square

func (s square) area() float64 {
	return s.side * s.side
}

// created interface: every type, that has method area() attached, can implement this interface

type shape interface {
	area() float64
}

// create new type

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// func to print Info

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	a := square{10}
	b := circle{5}
	fmt.Println(a.area()) //don't need interface here
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	info(a)
	info(b)
}
