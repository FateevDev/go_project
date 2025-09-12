package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	printArea() float64
}

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

func main() {
	t := Triangle{
		height: 10,
		base:   3,
	}

	s := Square{sideLength: 3}

	printArea(t)
	printArea(s)

}

func (t Triangle) printArea() float64 {
	return 0.5 * t.height * t.base
}

func (s Square) printArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(s Shape) {
	fmt.Println("shape", reflect.TypeOf(s).Name(), "has area: ", s.printArea())
}
