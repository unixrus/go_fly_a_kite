/*
functions
- passing values of diff types
- slice
- static global
- pointers
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

// function for calculatn average of int items in slice
func averageOf(tempsli []int) int { // <- funcions signature
	totalSum := 0
	for _, itemvalue := range tempsli {
		totalSum += itemvalue
	}
	return totalSum / len(tempsli)
	// return totalSum / float64(len(tempsli)) to correct divisor tye
}

// global paramter outside any functions
var param1 int = 1
var param2 int = 2

// func params to add global vars
func addp1p2them() int {

	return param1 + param2
}

// add with fixed # of params
func add3intItems(a, b, c int) int {
	return a + b + c
}

// pointer params for functions
func addmyPtrs(a, b *int) int {
	return *a + *b
}

// struct type
type Cube struct {
	depth  float64
	width  float64
	height float64
}

// method bound to Cube struct above
func (c *Cube) volume() float64 {
	return float64(c.depth) * float64(c.width) * float64(c.height)
}

func main() {
	fmt.Println("Action ")
	data := []int{3, 4, 6, 7, 8}
	fmt.Println(averageOf(data))

	fmt.Println(addp1p2them())

	fmt.Println(add3intItems(2, 3, 4))

	x := 6
	y := 7
	fmt.Println(addmyPtrs(&x, &y)) // param needs addrs of params

	dataC := Cube{2, 2.7, 2}
	fmt.Println(dataC.volume())
}

/*
Action

*/
