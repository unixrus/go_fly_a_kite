/*
	 operators
		arithmatic
		relationl operators
		logical operator
		bitise operators
		address of operator (&)
		pointer indirection(*) operator
*/
package main

import "fmt"

/* basic string ops */
func main() {
	x := 7
	y := 3
	// arithmatic
	fmt.Println(x+y, x-y, x*y, x/y, x%y)
	// 10 4 21 2 1

	// relationl operators
	fmt.Println(x == y, x != y, x > y, x < y, x >= y)
	//          false   true    true   false  true

	// logical operator
	fmt.Println(x > y && x >= y) // and  yields true
	fmt.Println(x > y || x < y)  // or   yields true
	fmt.Println(!true, !false)   // not  yields false true

	// bitwise operator
	fmt.Println(2 << 1) // 0010 -> 0100 = 4 i.e. bit shift 2 one place to left
	// other operators << >> & | ^ &^

	// assignmnt operator  i.e.  var = value
	// both side must be same type
	//  -= += *= /=

	// pointer operators

	x1 := "hello"  // short assign var x to helolo
	addrOfx := &x1 // address of operator x
	fmt.Println("value of x1 ", x1)
	fmt.Println("mem location addrOfx ", addrOfx)  // mem loc of x1 operator
	fmt.Println("contents of location ", *addrOfx) // pointer indirection opertor
	*addrOfx = "bye"                               // indirection operator (*) to change contents of mem loc
	fmt.Println("value of x1 ", x1)
	/*
		value of x1  hello
		mem location addrOfx  0xc0000120a0
		contents of location  hello
		value of x1  bye
	*/
}
