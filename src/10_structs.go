/*
structs type
*/
package main

import "fmt"

// exportable object called Box
type Box struct {
	depth  float64
	width  float64
	height float64
}

func main() {

	b := Box{3.0, 4.2, 5.3}
	fmt.Println(b) // {3 4.2 5.3}

	// struct literals can used to define in any ordrer using field names
	c := Box{depth: 3.0, height: 4.2, width: 5.3}

	fmt.Println(c) // {3 5.3 4.2}

	c.width = 7.9
	fmt.Println(c) // {3 7.9 4.2}

	// chnge attribute of Box via pointers
	ptr := &c
	(*ptr).depth = 6
	fmt.Println(c) // {6 7.9 4.2}

	ptr.depth = 9
	fmt.Println(c) // {6 7.9 4.2}
}
