package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "channel 1 response!"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("channel 1 timeout!")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "channel 2 response!"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
