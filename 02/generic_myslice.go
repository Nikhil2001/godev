package main

import "fmt"

type sliceInt []int

func (s sliceInt) Len() int {
	return len(s)
}

type sliceFloat []float64

func (s sliceFloat) Len() int {
	return len(s)
}

type slice[T any] []T

func (s slice[T]) Len() int {
	return len(s)
}

func main() {
	sInt := sliceInt{1, 2, 3}
	fmt.Println(sInt.Len())

	sFloat := sliceFloat{1.1, 2, 3}
	fmt.Println(sFloat.Len())

	slInt := slice[int]{1, 2, 3}
	fmt.Println(slInt.Len())
	slFloat := slice[float64]{1.1, 2, 3}
	fmt.Println(slFloat.Len())
	sString := slice[string]{"nik", "mara"}
	fmt.Println(sString.Len())
}
