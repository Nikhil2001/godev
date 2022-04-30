package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Println(b)

	str := "Programming in Go"
	b = []byte(str)
	fmt.Println(b)
	fmt.Printf("%s\n", b)
	fmt.Println(string(b))
	str = "世界"
	r := []rune(str)
	for _, v := range r {
		fmt.Println(v)
	}
	fmt.Printf("%d\n", r[0])

	fmt.Printf("%s\n", string(r))

}
