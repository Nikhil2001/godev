package main

import (
	"fmt"
	"sort"
)

func main() {

	x := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(x)
	sort.Ints(x)
	fmt.Println(x)
	sort.Sort(sort.Reverse(sort.IntSlice(x)))
	fmt.Println(x)

}
