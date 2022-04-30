package main

import "fmt"

func IsEqual[T comparable](x, y T) bool {
	return x == y
}

func main() {
	fmt.Println(IsEqual(1, 2))
	fmt.Println(IsEqual("nik", "nik"))

	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	fmt.Println(IsEqual(a, b))

}
