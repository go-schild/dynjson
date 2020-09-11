package main

import (
	"fmt"

	"github.com/go-schild/dynjson"
)

var jsonString = `[
	{"fname": "John", "lname": "Doe"},
	{"fname": "Jane", "mname": "Maria", "lname": "Dane"}
]`

func main() {
	list, _ := dynjson.ParseList(jsonString)

	for _, item := range list {
		obj := item.Object()

		if obj.Has("fname") {
			fmt.Println("First name:", obj.String("fname"))
		}
		if obj.Has("mname") {
			fmt.Println("Middle name:", obj.String("mname"))
		}
		if obj.Has("lname") {
			fmt.Println("Last name:", obj.String("lname"))
		}

		fmt.Println()
	}
}
