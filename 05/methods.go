package main

import (
	"fmt"
	"os"
	"strconv"
)

type M2x2 [2][2]int

func (a *M2x2) add(b M2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

func main() {

	if len(os.Args) != 9 {
		return
	}
	var k [8]int
	for i, v := range os.Args[1:] {
		i1, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		k[i] = i1
	}

	a := M2x2{{k[0], k[1]}, {k[2], k[3]}}
	b := M2x2{{k[4], k[5]}, {k[6], k[7]}}
	a.add(b)
	fmt.Println(a)

}
