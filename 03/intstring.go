package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Enter the args")
		return
	}

	arg := os.Args[1]

	n, err := strconv.ParseInt(arg, 10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	strn := strconv.Itoa(int(n))
	fmt.Println(strn)

	str := strconv.FormatInt(n, 2)
	fmt.Println(str)

	fmt.Printf("%f\n", 2.3)
	val, err := strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
