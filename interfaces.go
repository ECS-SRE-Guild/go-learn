package main

import (
	"fmt"
	"math"
)

// I have changed the names of the types and interface methods because they conflict with
// the go native library

type geometry interface {
	areaa() float64
	perimm() float64
}

type recta struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r recta) areaa() float64 {
	return r.width * r.height
}
func (r recta) perimm() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) areaa() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimm() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.areaa())
	fmt.Println(g.perimm())

}

func main() {
	r := recta{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
