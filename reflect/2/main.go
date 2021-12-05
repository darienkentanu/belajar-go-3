package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{"john wick", 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		// reflect.ValueOf("wick"),
		reflect.ValueOf("darien"),
	})

	fmt.Println("nama :", s1.Name)
}
