package main

import (
	"fmt"
	"log"
)

func main() {
	log.Fatalln(1, 2, 4)
	fmt.Println("this will not be printed")
}
