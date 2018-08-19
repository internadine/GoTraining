package main

import "fmt"

func main() {
	age := 40
	fmt.Println(&age) // 0xc042054058
	fmt.Println("************")

	changeMe(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 20

}
