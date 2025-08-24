package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 25}
	val := reflect.ValueOf(p)
	for i := 0; i < val.NumField(); i++ {
		fmt.Println("Field", i, ":", val.Field(i))
	}
}
