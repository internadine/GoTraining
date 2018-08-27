package main

import "fmt"

type person struct {
	First string
	Age   int
}

func main() {
	p1 := &person{"James", 30}

	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.First, " - ", p1.Age)
}
