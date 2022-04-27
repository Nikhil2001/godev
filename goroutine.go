package main

import (
	"fmt"
	"time"
)

func Print(i int) {
	time.Sleep(time.Second)
	fmt.Println(i)
}

func main() {
	//start := time.Now().Local().Second()
	for i := 1; i <= 5; i++ {
		go Print(i) // sequemtially would have taken more than 5 secs
	}

	time.Sleep(2 * time.Second)
	//fmt.Println("duration", time.Now().Local().Second()-start)

}
