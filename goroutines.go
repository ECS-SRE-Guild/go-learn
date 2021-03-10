package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct FIRST")

	go f("goroutine one SECOND")
	go f("goroutine two THIRD")

	go func(msg string) {
		fmt.Println(msg + " <IN MAIN> ")
		go f("goroutine three <IN MAIN> FOURTH")
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
