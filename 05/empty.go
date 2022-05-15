package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1 float64
	F2 S1
}

func Print(s any) {
	fmt.Println(s)
}

func main() {

	s1 := S1{10, "go"}
	s2 := S2{F1: 2, F2: s1}
	Print(s1)
	Print(s2)
	Print(123)
	Print(2 + 3i)
}
