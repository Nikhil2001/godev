package main

import "fmt"

func returnNumber() interface{} {
	return 12
}

func main() {
	a := returnNumber()
	num, ok := a.(int)
	num++
	fmt.Println(num, ok)
}
