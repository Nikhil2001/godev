package main

import "fmt"

func returnValue() any {
	return 12
}

type S1 struct {
	F1 int
}

func main() {
	i := returnValue()
	v, ok := i.(S1)
	fmt.Println(v, ok)

	v1, ok := i.(int64)
	v1++
	fmt.Println(v1, ok)

	v2, ok := i.(int)
	v2++
	fmt.Println(v2, ok)

	// _ = i.(bool) panic

}
