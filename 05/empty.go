package main

import "fmt"

type S1 struct {
	F1 string
	F2 int
}

type S2 struct {
	F1 int
	F2 float64
}

func Print(i interface{}) {
	fmt.Println(i)
}

func main() {
	var s1 = S1{F1: "field1", F2: 1}
	var s2 = S2{F1: 1, F2: 2.0}
	fmt.Println(s1)
	fmt.Println(s2)
	Print(s1)
	Print(s2)
	//defer func() { fmt.Println(recover()) }()
	_, _ = interface{}(s1).(int)

}
