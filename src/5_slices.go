/*
slices
- passed by reference
- so more efficient to pass slice rather than array to function
i.e. arrays make an entire copy each time
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

func main() {

	// zero slice
	var myslice []int
	fmt.Println("zero slice ", myslice)
	fmt.Println("zero slice is nil? ", myslice == nil)
	// zero slice  []
	// zero slice is nil?  true

	// make non zero slice with make(type, len, capacity)
	myslice1 := make([]int, 4, 4) // make slice of ints, with 4 length, capcity optional
	fmt.Println("slice ", myslice1)
	fmt.Println("this slice is nil? ", myslice1 == nil)
	// slice  [0 0 0 0]
	// this slice is nil?  false

	myslice2 := []int{5, 6, 7, 8, 9}
	fmt.Println("slice ", myslice2)
	fmt.Println("this slice is nil? ", myslice2 == nil)
	// slice  [5 6 7 8 9]
	// this slice is nil?  false

	myslice2[0] = 50
	fmt.Println("// slice ", myslice2)
	fmt.Println("// this slice is nil? ", myslice2 == nil)
	fmt.Println("// length of this slice? ", len(myslice2))
	// slice  [50 6 7 8 9]
	// this slice is nil?  false
	// length of this slice?  5

	// append to slice
	fmt.Println("// orig slice ", myslice2)
	myslice2 = append(myslice2, 10, 11, 12)
	fmt.Println("// appended slice ", myslice2)
	// orig slice  [50 6 7 8 9]
	// appended slice  [50 6 7 8 9 10 11 12]

	// appending slices ... note elipses ...
	printSlice(myslice2)
	// len=8 cap=10 [50 6 7 8 9 10 11 12]
	myslice2 = append(myslice2, myslice1...)
	printSlice(myslice2)
	// len=12 cap=20 [50 6 7 8 9 10 11 12 0 0 0 0]

	// make copy of myslice2 as copyOfMyslice2
	// init new slce with same specs - type lens cap etc
	copyOfMyslice2 := make([]int, len(myslice2))
	// make copy from src myslice2
	copy(copyOfMyslice2, myslice2)
	printSlice(copyOfMyslice2)
	// len=12 cap=12 [50 6 7 8 9 10 11 12 0 0 0 0]

	myslice31 := copyOfMyslice2[5:10] // elemnts 5,6,7,8,9
	printSlice(myslice31)
	myslice32 := copyOfMyslice2[:6] // elements 0 to 6th element
	printSlice(myslice32)
	myslice33 := copyOfMyslice2[6:] // elements starting at index val of 6 thru end of slice
	printSlice(myslice33)
	// len=5 cap=7 [10 11 12 0 0]
	// len=6 cap=12 [50 6 7 8 9 10]
	// len=6 cap=6 [11 12 0 0 0 0]
}

func printSlice(sl []int) {
	fmt.Printf("// len=%d cap=%d %v\n", len(sl), cap(sl), sl)
}

/*
$ go run 5_slices.go
zero slice  []
zero slice is nil?  true
slice  [0 0 0 0]
this slice is nil?  false
slice  [5 6 7 8 9]
this slice is nil?  false
// slice  [50 6 7 8 9]
// this slice is nil?  false
// length of this slice?  5
// orig slice  [50 6 7 8 9]
// appended slice  [50 6 7 8 9 10 11 12]
// len=8 cap=10 [50 6 7 8 9 10 11 12]
// len=12 cap=20 [50 6 7 8 9 10 11 12 0 0 0 0]
// len=12 cap=12 [50 6 7 8 9 10 11 12 0 0 0 0]
// len=5 cap=7 [10 11 12 0 0]
// len=6 cap=12 [50 6 7 8 9 10]
// len=6 cap=6 [11 12 0 0 0 0]

more experiments --> https://go.dev/tour/moretypes/15

*/
