package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./file.csv", os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
	newLine := []string{"Yuzvi chahal", "India"}
	csvWriter := csv.NewWriter(f)
	err = csvWriter.Write(newLine)
	if err != nil {
		return
	}
	csvWriter.Flush()
	fmt.Println("sucess")
}
