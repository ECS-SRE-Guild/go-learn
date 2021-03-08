package main

import (
	"fmt"
	"strings"
)

func simpleping() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println(msg)
	if msg == "ping" {
		fmt.Println("pong")
	}
}

func snd(msgchan chan string, msgs []string) {
	for i := 0; i < len(msgs); i++ {
		msgchan <- msgs[i]
	}
}

func rcv(msgchan chan string) string {
	var sentence []string
	var msglen int = len(msgchan)
	for v := range msgchan {
		fmt.Println(len(sentence), sentence)
		if len(sentence) == msglen-1 { // when we reach next to last in the array, do the last.
			sentence = append(sentence, v)
			break
		} else {
			sentence = append(sentence, v)
		}
	}
	return strings.Join(sentence, " ")
}

func main() {
	simpleping()

	var words = []string{"buffered", "channel", "one", "two", "three", "four", "five", "six"}
	msgchan := make(chan string, len(words))
	snd(msgchan, words)
	fmt.Println(rcv(msgchan))

}
