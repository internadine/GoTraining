package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubeZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubeZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}
	p2 := doubeZero{
		person: person{
			First: "Nadine",
			Last:  "Wischmeier",
			Age:   40,
		},
		LicenseToKill: false,
	}
	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
}
