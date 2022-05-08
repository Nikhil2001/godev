package main

import "fmt"

func main() {
	var aMap map[int]string
	fmt.Println(aMap == nil, len(aMap))

	aMap = make(map[int]string)
	aMap[0] = "zero"
	fmt.Println(aMap, len(aMap))
	aMap[1] = "one"
	for k, v := range aMap {
		fmt.Println(k, v)
	}
}
