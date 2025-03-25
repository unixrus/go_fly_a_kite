/*
variadic functions
- multiple params
- returns
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

func add2nPrint(x, y int) {
	// no returns
	fmt.Println(x + y)
}

// func with multiple returns
func compareVars(x, y int) (a, b int, c bool) {
	a = x + y
	b = x - y
	c = x > y
	return // names not necessary here if they match in code.
}

// variadic func with variable # of params
// pack operator ...
func addAll(items ...int) ([]int, int) {
	sumofItems := 0
	for _, item := range items {
		sumofItems += item
	}
	return items, sumofItems
}

func main() {
	add2nPrint(2, 4) // print 6
	fmt.Println(compareVars(5, 7))
	fmt.Println(compareVars(9, 7))

	// unpack operator
	dataslice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(dataslice)
	// fmt.Println(dataslice...)

	// trailing ... to unpack dataslice
	fmt.Println(addAll(dataslice...))

	// hardcoded param list of 0 or more args
	// params are converted to new allocated slice
	// and passed function as variadic param
	fmt.Println(addAll(1, 2, 3, 4, 5, 6, 7, 8))
}

/*
$ go run 9b_func_returns.go
6
12 -2 false
16 2 true
[1 2 3 4 5 6]
[1 2 3 4 5 6] 21
[1 2 3 4 5 6 7 8] 36
*/
