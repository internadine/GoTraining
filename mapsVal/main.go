package main

import (
	"fmt"
)

func main() {
	myGreetings := map[int]string{
		0: "Good morning",
		1: "bonjour",
		2: "boenos diaz",
		3: "guten Morgen",
	}
	fmt.Println(myGreetings)
	if value, exists := myGreetings[4]; exists {
		fmt.Println(value)
		delete(myGreetings, 2)
	} else {
		fmt.Println(exists)
	}
	fmt.Println(myGreetings)
}
