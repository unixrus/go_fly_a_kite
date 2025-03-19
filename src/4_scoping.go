/*
vars in global and local scopes
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

func main() {
	// main x in scope
	x := 5
	fmt.Println("main x is ", x)

	for x := 0; x < 4; x++ {
		fmt.Println("loop x is ", x)
	}
	fmt.Println("main x is still ", x)
}

/*
main x is  5
loop x is  0
loop x is  1
loop x is  2
loop x is  3
main x is still  5
*/
