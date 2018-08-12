package main

import "fmt"

func main() {
	switch "Sunday" {
	case "Monday":
		fmt.Println("Start to work")
	case "Tuesday":
		fmt.Println("Second day work")
	case "Wednesday":
		fmt.Println("Mid of week")
	case "Thursday":
		fmt.Println("almost done")
	case "Friday":
		fmt.Println("done the week, yeah")
		fallthrough
	case "Saturday":
		fmt.Println("Stay at home")
	default:
		fmt.Println("finally it is sunday")
	}
}
