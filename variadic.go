package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func sentence(words ...string) {
	snt := ""
	for _, word := range words {
		snt += word + " "
	}
	fmt.Println(snt)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	sentence("This", "is", "my", "useless", "function")
}
