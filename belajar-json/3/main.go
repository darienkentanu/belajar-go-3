package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"`
	Age      int
}

func main() {
	object := []User{
		{"darien kentanu", 25},
		{"qwerty", 0},
	}

	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
	// fmt.Println(jsonData)
}
