package main

import "fmt"

func main() {

	fmt.Println("Enter your name below :")
	var name string
	fmt.Scanln(&name)
	fmt.Println("welcome ", name)
}
