package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i interface{} = 2
	val := reflect.ValueOf(i)
	fmt.Println(reflect.Int == val.Kind())

	var s interface{} = [4]int{1, 2, 3, 4}
	val1 := reflect.ValueOf(s)
	fmt.Println(reflect.Array == val1.Kind())

	for i := 0; i < val1.Len(); i++ {
		fmt.Println(val1.Index(i).Interface())
	}
}
