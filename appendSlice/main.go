package main

import "fmt"

func main() {
	fmt.Println("***************************************")
	x := make([]int, 0, 5)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("***************************************")
	for i := 0; i < 80; i++ {
		x = append(x, i)
		fmt.Println("Len:", len(x), "Capacity: ", cap(x), "Value: ", i)
	}

	y := make([]string, 3, 10)

	y[0] = "Sugar"
	y[1] = "Chocolate"
	y[2] = "Cake"
	y = append(y, "Candy")
	y = append(y, "Lolly")
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

}
