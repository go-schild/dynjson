package main

import (
	"fmt"

	"github.com/go-schild/dynjson"
)

func main() {
	obj := dynjson.NewJsonObject()

	obj.SetString("fname", "John")
	obj.SetString("nname", "Doe")

	// {"fname":"John","nname":"Doe"}
	fmt.Println(obj.ToString())
}
