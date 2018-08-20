package main

import "fmt"

func main() {
	var x [58]string
	fmt.Println(x)               //empty array
	fmt.Println(len(x))          // length of string
	for i := 65; i <= 122; i++ { // fill array with letters
		x[i-65] = string(i) // transform number in letter string(i)
	}
	fmt.Println(x)
	fmt.Println(x[42]) //show the 43th letter in the array
}
