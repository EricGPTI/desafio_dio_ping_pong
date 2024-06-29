package main

import (
	"ping_pong/ping"
	"ping_pong/pong"
	"fmt"
)


func main () {
	c := make(chan string)
	done := make(chan struct{})

	go ping.SenderPing(c, done)
	go pong.SenderPong(c, done)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
		done <- struct{}{}
		fmt.Println(<-c)
		done <- struct{}{}
	}
}