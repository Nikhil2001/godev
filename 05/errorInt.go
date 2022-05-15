package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type emptyFile struct {
	read  int
	ended bool
}

func (e emptyFile) Error() string {
	return fmt.Sprintf("Ended with io.EOF (%t),but read %d", e.ended, e.read)

}

func IsFileEmpty(e error) bool {
	val, ok := e.(emptyFile)
	if ok {
		if val.ended == false && val.read == 0 {
			return true
		}
	}

	return false

}

func readfile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var n int
	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			n += len(line)
		} else if err == io.EOF {
			return emptyFile{ended: true, read: n}
		} else {
			return err
		}
	}
}

func main() {
	err := readfile("file.txt")
	if IsFileEmpty(err) {
		fmt.Println("empty")
	} else {
		fmt.Println(err)
	}

}
