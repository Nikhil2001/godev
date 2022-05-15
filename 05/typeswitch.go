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

func Test(v any) {
	switch T := v.(type) {
	case S1:
		fmt.Println("struct S1", T)
	case S2:
		fmt.Println("struct S2", T)
	default:
		fmt.Printf("%T %v \n", T, T)
	}
}

func main() {
	Test(1)
	Test("go")
	s1 := S1{10, "go"}
	s2 := S2{F1: 2, F2: s1}
	Test(s1)
	Test(s2)
}
