package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m)
	fmt.Println("*******")
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	fmt.Println(z)
	z[0] = "Here we go!"

}
