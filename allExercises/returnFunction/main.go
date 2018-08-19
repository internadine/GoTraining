package main

import "fmt"

func half(z int) (int, bool) {
	return z / 2, z%2 == 0
}

func main() {
	fmt.Println(half(5))
}
