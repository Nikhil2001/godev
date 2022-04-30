package main

import "fmt"

type Myint int

type Numeric interface {
	~int | float64 | string | complex128
}

func add[T Numeric](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(3, 4))

	var a Myint = 2
	var b Myint = 20
	fmt.Println(add(a, b))

}
