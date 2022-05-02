package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return -1, fmt.Errorf("cannot divide by %b", b)
	}
	return a / b, nil
}

func main() {
	err1 := errors.New("error")
	err2 := fmt.Errorf("error %d", 2)
	fmt.Println(err1, err2)
	fmt.Println(err1.Error(), err2.Error())

	fmt.Println(div(2, 4))
	fmt.Println(div(2, 0))
}
