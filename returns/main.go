package main

import "fmt"

func main() {
	fmt.Println(greet("Nadine"))
}

func greet(name string) string {
	return fmt.Sprint(name)
}
