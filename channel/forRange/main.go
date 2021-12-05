package main

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	// fmt.Println(&ch)
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	messages := make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}
