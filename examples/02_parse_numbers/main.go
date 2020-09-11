package main

import (
	"fmt"

	"github.com/go-schild/dynjson"
)

func main() {
	jsonString := `{"a": 1, "b": 2, "pi": 3.14159265358979}`
	obj, _ := dynjson.ParseObject(jsonString)

	fmt.Println(obj.Int("a"))
	fmt.Println(obj.Int("b"))
	fmt.Println(obj.Float64("pi"))
	fmt.Println(obj.Int("pi"))
}
