package main

import "fmt"

func main() {
	myGreetings := map[int]string{
		0: "Guten Morgen",
		1: "Guten Tag",
		2: "Guten Abend",
		3: "Guten Nacht",
	}
	fmt.Println(myGreetings)

	for key, val := range myGreetings {
		fmt.Println(key, " - ", val)
	}
}
