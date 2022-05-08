package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const CSVFILE = "./contacts.txt"

type Entry struct {
	Name    string
	SurName string
	Number  int
}

var list = []Entry{}

func List() {
	for _, v := range list {
		fmt.Println(v)
	}
}

func Delete(entry Entry) error {
	for i, v := range list {
		if v.Name == entry.Name && v.SurName == entry.SurName && v.Number == entry.Number {
			list = append(list[:i], list[i+1:]...)
			list = list[:len(list)-1]
			return nil
		}
	}
	return errors.New("contact not found,cannot delete")
}

func ConvertToEntry(entry string) (Entry, error) {
	fields := strings.Split(entry, ",")
	num, err := strconv.Atoi(fields[2])
	if err != nil {
		return Entry{}, errors.New("Invalid Phone Number")
	}
	return Entry{Name: fields[0], SurName: fields[1], Number: num}, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: Insert | Delete | Modify | Search | List <arguments>")
		return
	}

	option := os.Args[1]
	entry, err := ConvertToEntry(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	switch option {
	case "Insert":
		list = append(list, entry)
		fmt.Println("new contact added.")
	case "Delete":
		err := Delete(entry)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("contact deleted.")
	case "List":
		fmt.Println("Showing all contacts: ")
		List()
	case "Modify":

		fmt.Println("contact Modified.")
	default:
		fmt.Println("Invalid Option")
		fmt.Println("available : Insert | Delete | Modify | Search | List")
	}
}
