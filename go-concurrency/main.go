package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)

	// send data to channel in right side of arrow
	// true because type of bool in chan
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	// send data to channel in right side of arrow
	// true because type of bool in chan
	// you can send multiple values from different go routines through the same channel, but if 1 function complete the execution is done
	doneChan <- true
	// close it because we know that channel is done in slowGreet (its only make sense because we know)
	close(doneChan)
}

func main() {
	// Channel = value that can be used as communication when working with go routines
	// make(chan + type of data that send through channel)
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// go means tells go to run parallel (go routines)
	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	// reading from done channel
	// for _, done := range dones {
	// 	<-done
	// }

	// for loop direct to done channel
	// error because no value left in channel
	for range done {
	}
}
