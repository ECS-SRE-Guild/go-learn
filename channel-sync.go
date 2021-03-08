package main

import (
	"fmt"
	"time"
)

func signal() chan bool {
	return make(chan bool, 1)
}

func worker(done chan bool, wkrnum int) {
	fmt.Println("BEGIN: ", wkrnum)
	fmt.Println("Wait 5 seconds...", wkrnum)
	time.Sleep(time.Duration(5 * time.Second))
	fmt.Printf("%d plus %d equals %d\n", wkrnum, 5, 5+wkrnum)
	time.Sleep(time.Duration(5000))
	fmt.Println("END: ", wkrnum)

	done <- true
}

func main() {

	done1 := signal()
	go worker(done1, 1)
	<-done1
	done2 := signal()
	go worker(done2, 2)
	<-done2
	done3 := signal()
	go worker(done3, 3)
	<-done3
}
