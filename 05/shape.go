package main

import (
	"fmt"
	"math"
)

type shape interface {
	Perimeter() float64
}

type circle struct {
	radius float64
}

func (c *circle) Perimeter() float64 {
	return c.radius * math.Pi * 2
}

func main() {
	r := 7.0
	var c shape = &circle{r}
	fmt.Println(c.Perimeter())
}
