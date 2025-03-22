/*
flow control with interating over collections with range
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

func main() {
	mysli := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(mysli); i++ {
		fmt.Println(i, mysli[i])
	}

	// same as above using range
	mysli1 := []int{1, 2, 3, 4, 5}
	for i, value := range mysli1 {
		fmt.Println(i, value)
	}

	mymap := map[string]int{
		"item1": 15,
		"item2": 25,
		"item3": 35,
	}
	for key, value := range mymap {
		fmt.Println(key, value)
	}
	for key := range mymap { // key is required
		fmt.Println(key)
	}
	for _, value := range mymap { // key is required
		fmt.Println(value)
	}
}

/*
Action

*/
