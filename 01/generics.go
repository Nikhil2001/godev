package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	Print([]int{1, 2, 3, 4, 5})
	Print([]string{"one", "two", "three", "four", "five"})

}
