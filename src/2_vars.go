package main

import "fmt"

/* basic string ops */
func main() {
	var msg string = "hello world"
	fmt.Println(msg)

	var msg1 string
	msg1 = "hello world1"
	fmt.Println(msg1)

	msg2 := "infer the var type"
	fmt.Println(msg2)

	msg3 := 1
	fmt.Println(msg3)

	// show zero values for each type
	var x int
	var truth bool
	var thetext string
	var pointer *string
	fmt.Println("integer ", x)
	fmt.Println("   bool ", truth)
	fmt.Println(" string ", thetext)
	fmt.Println("pointer ", pointer)
	/* results
		integer  0
	       bool  false
	     string
	    pointer  <nil>
	*/

	//  constant vars that are fixed
	const pi = 3.14
	// pi = 3.1234    fails as pi cannot be updated
	fmt.Println(pi)

	// constant is an integer but allows float addition
	const myconst = 1
	fmt.Printf("type: %T, value: %v \n", myconst+1.1, myconst+1.1)
	// type: float64, value: 2.1

	// here the myconst has to be converted to float for addition
	const myconst1 int = 1
	fmt.Printf("type: %T, value: %v \n", float64(myconst1)+1.1, float64(myconst1)+1.1)
}
