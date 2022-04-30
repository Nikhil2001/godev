package main

import "fmt"

func main() {

	byteSlice := make([]byte, 12)
	fmt.Println(byteSlice)

	byteSlice = []byte("I am a Programmer")
	fmt.Println(byteSlice)
	fmt.Println(string(byteSlice))

	fmt.Printf("%s\n", byteSlice)

}
