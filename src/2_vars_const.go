/*
create simple function
func are run at compile time
const vars
convert int to float64
*/
package main

import "fmt"

// iota assign incrementing values starting at 0
const (
	Sun = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

/* basic const var ops */
func main() {

	// constant is an integer but allows float addition
	const myconst = 1
	fmt.Printf("type: %T, value: %v \n", myconst+1.1, myconst+1.1)
	// type: float64, value: 2.1

	// here the myconst1 has to be converted to float for addition
	const myconst1 int = 1
	fmt.Printf("type: %T, value: %v \n", float64(myconst1)+1.1, float64(myconst1)+1.1)

	// myconst2 cannot be defined as it cannot be determined at compile time
	// const myconst2 = returnInt()
	// const myconst2 int = returnInt()    will also not work
	// fmt.Printf("type: %T, value: %v \n", float64(myconst2)+1.1, float64(myconst2)+1.1)

	fmt.Println(Sun, Mon, Tue, Wed, Thu, Fri, Sat)
	// print 0 1 2 3 4 5 6
}

// return value is an int
func returnInt() int {
	return 1
}
