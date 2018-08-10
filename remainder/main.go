package main

import "fmt"

func main() {
	for i := 5; i < 30; i++ {
		if i%2 == 1 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}
	}
}
