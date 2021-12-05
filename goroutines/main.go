package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// fmt.Println("num cpu =", runtime.NumCPU())
	runtime.GOMAXPROCS(2)

	go print(25, "halo")
	print(25, "apa kabar")

	var input string
	fmt.Scanln(&input)
}
