package main

import "fmt"

func main() {

	s1 := []int{1}
	s2 := []int{2, 9, 4, 5}
	s3 := []int{2, 3}
	//copy(dest,src)
	fmt.Println(s3, s1)
	copy(s3, s1)
	fmt.Println(s3, s1)

	fmt.Println(s2, s3)
	copy(s2, s3)
	fmt.Println(s2, s3)
}
