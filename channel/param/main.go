package main

import (
	"fmt"
	"runtime"
)

func printMessage(what <-chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var messages = make(chan string)

	for _, each := range []string{"darien", "fitra", "meilina"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
