package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPtoS() *Person {
	return &Person{}
}

func newS() Person {
	return Person{}
}

func newPerson(name string, age int) *Person {
	if len(name) == 0 {
		return &Person{name: "Unknown", age: age}
	}
	return &Person{name, age}
}

func main() {
	s1, s2, s3, s4 := newPtoS(), newS(), newPerson("Nikhil", 28), newPerson("", 30)
	fmt.Println(s1, s2, s3, s4)
	fmt.Println(s3.name) // . operator works for pointers too

}
