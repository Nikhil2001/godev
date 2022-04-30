package main

import (
	"fmt"
	"time"
)

func main() {
	v, _ := time.Parse(time.Kitchen, "01:03PM")
	fmt.Println(v)
}
