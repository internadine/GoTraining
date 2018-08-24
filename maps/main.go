package main

import "fmt"

func main() {
	var greetings = make(map[string]string, 5)
	greetings["Tim"] = "Good morning"
	greetings["Jenny"] = "Bonjour"

	fmt.Println(greetings)
}
