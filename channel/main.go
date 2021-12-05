package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("darien")
	go sayHelloTo("fitra")
	go sayHelloTo("meilina")

	var message1 = <-messages
	var message2 = <-messages
	var message3 = <-messages

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)
}
