package main

import "fmt"

type Astruct struct {
	a int
	b int
}

func main() {

	var i int = 2
	fmt.Println(i, &i)

	p := &i
	fmt.Println(p)
	fmt.Println(&p)

	*p = 22
	fmt.Println(i)

	var str *Astruct
	fmt.Printf("%+v %v\n", str, str == nil)

	var str1 = new(Astruct)
	fmt.Printf("%+v\n", str1)
}
