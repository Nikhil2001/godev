package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Perimeter() float64
}

type Circle struct {
	r float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func Assert(i any) {
	c, ok := i.(Shape)
	fmt.Printf("%T,%T\n", c, ok)

}

func main() {
	c1 := Circle{r: 2.0}
	Assert(c1)

	s1 := Shape(c1)

	c1, ok := s1.(Circle)
	fmt.Printf("%T,%T\n", c1, ok)

}
