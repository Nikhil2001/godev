package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(aSlice, cap(aSlice))
	i := 2
	bSlice := append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println(aSlice, cap(aSlice))
	fmt.Println(bSlice, cap(bSlice))

	bSlice = append(bSlice, 7)
	fmt.Println(aSlice, cap(aSlice))
	fmt.Println(aSlice, cap(aSlice))
}
