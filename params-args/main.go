package main

import "fmt"

func main() {
	greet("Jane", "Doe") //Jane is an argument
	greet("John", "Thunder")
}

func greet(fname, lname string) { //name is an parameter
	fmt.Println(fname, lname)
}
