/*
Errors - defer, panic, recover

setup a defer func recover to catch any panic and process

visual studio tips;
ctrl `  to toggle between code and terminal
select lines and ctrl / to comment out line with //
*/
package main

import (
	"fmt"
)

// fnction that will have a panic condition
func p(panicme string) {
	// trminate normal execution flow and issue a panic
	if panicme == "yes" {
		panic("Panic'ing as it is a definite impossible condition")
	} else if panicme == "maybe" {
		panic("Panic'ing as it is a impossible maybe condition")
	} else {
		fmt.Println("normal execution")

	}
}

// function that will handle panic condition
// using recover
func r() {
	err := recover()
	if err != nil {
		fmt.Println("print panic mesg: ", err)
	} else {
		fmt.Println("all good, cleanup and terminate")
	}

}

func real_condition() {
	x := [4]int{0, 1, 2, 3}
	// create anic condition by llegal access to array
	i := 5
	// x[i] = 11
	fmt.Println(x[i])
}

func main() {

	// setup defer with r func to execute after p func
	defer r()

	// just a sample harcoded panic call
	// p("yes")
	// print panic mesg:  Panic'ing as it is a definite impossible condition

	// p("maybe")
	// print panic mesg:  Panic'ing as it is a impossible maybe condition

	// p("no not really")
	// normal execution
	// all good, cleanup and terminate

	// trigger an actual runtime panic
	real_condition()
	// print panic mesg:  runtime error: index out of range [5] with length 4

}
