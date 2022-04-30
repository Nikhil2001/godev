package main

import "fmt"

func main() {
	slice1 := make([]int, 4)
	fmt.Println(slice1, cap(slice1), len(slice1))

	slice2 := make([]int, 4, 5)
	fmt.Println(slice2, cap(slice2), len(slice2))

	aslice := make([]int, 4)
	fmt.Println(aslice, cap(aslice), len(aslice))

	aslice = append(aslice, 3)
	fmt.Println(aslice, cap(aslice), len(aslice))

	aslice = append(aslice, []int{1, 2, 3, 5}...)
	fmt.Println(aslice, cap(aslice), len(aslice))

}
