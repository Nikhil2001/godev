package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

const (
	e = 0
	_
	f
	g = iota
	h
)

const (
	i = iota
	_
	j
	_
	k
	_
	_
	l
)

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)
	fmt.Println(i, j, k, l)

}
