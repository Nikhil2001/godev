package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Printf("type of os.Args: %T\n", os.Args)
	fmt.Printf("value of os.Args: %+v\n", os.Args)
	fmt.Println("length of os.Args:", len(os.Args))

	if len(os.Args) <= 1 {
		fmt.Println("Enter args please!")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		intValue, err := strconv.ParseInt(os.Args[i], 10, 0)
		if err != nil {
			fmt.Println(i, "arg is", os.Args[i], "cannot be string error:", err)
			continue
		}
		fmt.Println(i, "argument is", intValue)
	}
	fmt.Println("Program finished!")
}
