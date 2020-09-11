package main

import (
	"fmt"

	"github.com/go-schild/dynjson"
)

func main() {
	jsonString := `{"outer": {"inner": 5}}`
	obj, _ := dynjson.ParseObject(jsonString)

	outer := obj.Object("outer")
	fmt.Println(outer.Int("inner"))
}
