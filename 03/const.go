package main

import "fmt"

const (
	a = 1
	b
	c = iota
	_
	d
)

const (
	e = 2
	f
	g = f << iota
	h
)

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)

	const s1 = 123
	var f1 float64 = s1 * 2
	fmt.Println(f1)

}
