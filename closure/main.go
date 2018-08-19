package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		y := "What the hell"
		fmt.Println(y)
	}
	y := x
	fmt.Println(y)
}
