package main

import "fmt"

func main() {

	half := func(z int) (float64, bool) {
		return float64(z) / 2, z%2 == 0
	}
	fmt.Println(half(5))
}
