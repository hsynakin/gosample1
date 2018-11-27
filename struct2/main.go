package main

import (
	"fmt"
)

type myStruct struct {
	myMap map[string]string
}

func main() {
	foo := newMyStruct()
	foo.myMap["bar"] = "baz"
	fmt.Println(foo)

}

func newMyStruct() *myStruct {
	result := myStruct{}
	result.myMap = map[string]string{}
	return &result

}
