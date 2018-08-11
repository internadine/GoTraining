package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("see this", i)
	}
	nested()
}

func nested() {
	for i := 1; i < 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(i, "-", j)
		}
	}
}
