package main

import (
	"fmt"

	"github.com/go-schild/dynjson"
)

func main() {
	obj := dynjson.NewJsonObject()

	// JSON doesn't have different types of numbers (e.g. int or float32).
	// It only knows the type "number" for numbers.
	obj.SetNumber("a", 1)
	obj.SetNumber("b", 2)
	obj.SetNumber("pi", 3.14159265358979)

	// {"a":1,"b":2,"pi":3.14159265358979}
	fmt.Println(obj.ToString())
}
