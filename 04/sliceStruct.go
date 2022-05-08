package main

import (
	"fmt"
	"strconv"
)

type Item struct {
	name  string
	price int
}

func main() {

	list := []Item{}

	var name string
	var total int

	for i := 0; i < 10; i++ {
		name = "text" + strconv.Itoa(i)
		item := Item{name: name, price: i}
		list = append(list, item)
	}

	for _, i := range list {
		total += i.price
	}
	fmt.Println(total)

	for _, i := range list {
		fmt.Println(i)

	}

}
