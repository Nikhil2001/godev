package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("enter file")
		return
	}

	file := os.Args[1]
	path := os.Getenv("PATH")
	var paths = filepath.SplitList(path)

	for _, p := range paths {
		f := filepath.Join(p, file)
		fileinfo, err := os.Stat(f)
		if err == nil {
			if fileinfo.Mode().IsRegular() {
				if fileinfo.Mode()&0111 != 0 {
					fmt.Println(p)
				}
			}
		}
	}

}
