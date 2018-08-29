package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)
	fmt.Println(sort.StringsAreSorted(s))
	fmt.Println(s)
	i := []int{1, 2, 3, 7, 33, 556, 88}
	sort.Ints(i)
	fmt.Println(i)
	fmt.Println(sort.IntsAreSorted(i))

}
