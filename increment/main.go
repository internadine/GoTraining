package main

import (
	"fmt"
)

var x int

func increment() int {
	x++
	return x
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(increment())
	}

}
