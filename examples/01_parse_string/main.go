package main

import (
	"fmt"

	"github.com/go-schild/dynjson"
)

func main() {
	jsonString := `{"fname": "John", "lname": "Doe"}`
	obj, _ := dynjson.ParseObject(jsonString)

	if obj.Has("fname") {
		fmt.Println("First name:", obj.String("fname"))
	}
	if obj.Has("mname") {
		fmt.Println("Middle name:", obj.String("mname"))
	}
	if obj.Has("lname") {
		fmt.Println("Last name:", obj.String("lname"))
	}
}
