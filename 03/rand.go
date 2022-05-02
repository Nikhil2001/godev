package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {

	b := make([]byte, 6)
	rand.Read(b)
	fmt.Println(b)

	s := base64.URLEncoding.EncodeToString(b)
	fmt.Println(s)
}
