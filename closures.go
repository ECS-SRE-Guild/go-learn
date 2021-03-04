package main

import (
	"fmt"
	"math/rand"
	"time"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func words() func() string {
	wordlist := [11]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six",
		7: "seven", 8: "eight", 9: "nine", 10: "ten"}

	return func() string {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 10
		return wordlist[rand.Intn(max-min+1)+min]
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	newWords := words()
	fmt.Println(newWords())

}
