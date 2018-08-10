package main

import "fmt"

func main() {
	a := 42
	fmt.Println(&a) //prints 0xc420098000
	fmt.Println(a)  //prints 42

	var b *int = &a
	fmt.Println(b)  //prints 0xc420098000
	fmt.Println(*b) // prints 42 = this is called dereferencing

	*b = 10
	fmt.Println(a) //prints out 10
}
