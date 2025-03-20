/*
arrays
- index start at 0 for first elemtn like py
- fixed in length at initialization
- fixed type of elements
- assigning arrays to another var create a copy
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

func main() {
	var arr [4]int
	fmt.Println("Array ", arr)
	// Array [0 0 0 0 ]

	// using := short assignment operator
	arr1 := [4]int{0, 1, 2, 3}
	fmt.Println("Array ", arr1)
	// Array [0 1 2 3]

	// update element
	arr1[3] = 30
	fmt.Println("Array ", arr1)
	// Array  [0 1 2 30]
	fmt.Println("Array  length", len(arr1))
	// Array  length 4

	// two dinmensional array
	var arrMultiDimensional [3][2]int
	fmt.Println("multi dimention arr ", arrMultiDimensional)
	// multi dimention arr  [[0 0] [0 0] [0 0]]

	arrMD2 := [3][4]int{
		{0, 1, 2, 3},   /*  initializers for row indexed by 0 */
		{4, 5, 6, 7},   /*  initializers for row indexed by 1 */
		{8, 9, 10, 11}, /*  initializers for row indexed by 2 */
	}
	fmt.Println("multi dimention arr ", arrMD2)
	// multi dimention arr  [[0 1 2 3] [4 5 6 7] [8 9 10 11]]

}

/*
Action

*/
