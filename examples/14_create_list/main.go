package main

import (
	"fmt"
	"github.com/go-schild/dynjson"
)

func main() {
	list := dynjson.NewJsonList(nil)

	obj1 := dynjson.NewJsonObject()
	obj1.SetString("fname", "John")
	obj1.SetString("nname", "Doe")

	obj2 := dynjson.NewJsonObject()
	obj2.SetString("fname", "Jane")
	obj2.SetString("mname", "Maria")
	obj2.SetString("nname", "Dane")

	list.Append(obj1, obj2)
	fmt.Println(list.ToString())  // TODO doesn't work properly yet
}
