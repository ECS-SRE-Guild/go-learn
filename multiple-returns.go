package main

import (
	"fmt"
	"strings"
)

func vals() (int, int) {
	return 3, 7
}

func strs(a string) (string, string, string) {
	b := strings.Split(a, " ")
	return b[0], b[1], b[2]
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	x, y, z := strs("One Two Three")
	fmt.Println(x, y, z)

}
