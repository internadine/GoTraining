package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type secretAgend struct {
	person
	licenceToKill bool
}

func (p person) speak() string {
	return "Hello, I am not a super hero!"
}

func (s secretAgend) speak() string {
	return "Hello, Yes I am a super hero!"
}

type human interface {
	speak() string
}

func main() {
	p1 := person{"Nadine", "Wischmeier", 40}
	sh := secretAgend{
		person: person{
			"James",
			"Bond",
			50,
		},
		licenceToKill: true,
	}
	fmt.Println(p1.fname)
	fmt.Println(p1.speak())
	fmt.Println(sh.person.fname)
	fmt.Println(sh.speak())
}
