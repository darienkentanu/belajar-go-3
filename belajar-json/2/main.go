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
	var jsonString = `[
		{"Name": "darien kentanu", "Age": 25},
		{"Name": "qwerty", "Age": 0}
		]`

	var data []User
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data[0].Fullname)
	fmt.Println("user 2:", data[1].Fullname)
}
