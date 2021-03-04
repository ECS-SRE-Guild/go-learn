package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(fact(129))
	fmt.Println(fact(11))
	fmt.Println(fact(1))
	fmt.Println(fact(9321))
}
