package main

import "fmt"

type storage[T any] struct {
	cap []T
}

func (s storage[T]) Print() {
	for _, v := range s.cap {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	store := storage[int]{cap: []int{1, 2, 3}}
	store.Print()

	storeStr := storage[string]{cap: []string{"a", "aa", "aaa"}}
	storeStr.Print()
}
