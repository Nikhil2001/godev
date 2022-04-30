package main

import "fmt"

func main() {
	var a = []int{}
	fmt.Println(a, len(a), cap(a), a == nil)
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))

	var twoD [][]int
	twoD = append(twoD, []int{1})
	twoD = append(twoD, []int{2, 3})
	twoD = append(twoD, []int{4, 5, 6})

	for _, v := range twoD {
		fmt.Println(v)
	}
}
