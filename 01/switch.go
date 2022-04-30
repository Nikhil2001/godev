package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Provide Argument,Exiting...")
		return
	}

	switch os.Args[1] {
	case "0":
		fmt.Println("zero")
	case "1":
		fmt.Println("one")
	case "2", "3", "4":
		fallthrough
	default:
		fmt.Println("value is", os.Args[1])
	}

	val, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	switch {
	case val == 0:
		fmt.Println("value:", 0)
	case val == 1:
		fmt.Println("value:", 1)
	case val == 2, val == 3, val == 4:
		fallthrough
	default:
		fmt.Println("value:", val)
	}
}
