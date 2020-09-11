package main

import (
	"fmt"

	"github.com/go-schild/dynjson"
)

func main() {
	obj := dynjson.NewJsonObject()
	outer := dynjson.NewJsonObject()

	outer.SetNumber("inner", 5)
	obj.SetObject("outer", outer)

	// {"outer":{"inner":5}}
	fmt.Println(obj.ToString())
}
