package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	str, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	filepath := filepath.Join(str, "/log.txt")
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	log := log.New(f, "=> ", log.LstdFlags|log.Lshortfile)
	log.Println("nikhil")
	log.Println("reddy")
}
