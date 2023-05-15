package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("Hello world!"), write("I'm a Go developer"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(firstInputChannel, secondInputChannel <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func () {
		for {
			select {
			case message := <-firstInputChannel:
				outputChannel <- message
			case message := <-secondInputChannel:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	go func ()  {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(r.Intn(2000)))
		}
	}()

	return channel
}