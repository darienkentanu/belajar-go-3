package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name": "Darien Kentanu", "Age": 25}`
	var jsonData = []byte(jsonString)
	// var data1 map[string]interface{}
	// json.Unmarshal(jsonData, &data1)

	// fmt.Println("user :", data1["Name"])
	// fmt.Println("age  :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2
	// fmt.Println("user :", decodedData["Name"])
	// fmt.Println("age  :", decodedData["Age"])
	fmt.Println(decodedData)
}
