package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	obj := reflect.TypeOf(x)
	fmt.Println("\n", obj, obj.Name(), obj.Kind())
}

func main() {
	fmt.Println("111")
	reflectType([]string{"123", "321"})
}
