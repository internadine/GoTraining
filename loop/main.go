package main

import "fmt"

func main() {
	for i := 0; i < 50; i++ {
		fmt.Printf("%d - %b \n", i, i)
	}
}
