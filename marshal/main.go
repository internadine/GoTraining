package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{"James", "Bond", 40}
	fmt.Println(p1)
	fmt.Println(p1.First)
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))

}
