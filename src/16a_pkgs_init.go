package main

// init func before main.
// inits in imports go before these inits

import (
	"fmt"
	_ "strings"
)

// global var env declared
var env string

func init() {
	fmt.Println("step #1 init")
}
func init() {
	fmt.Println("step #2 init")
}
func init() {
	env = "DEVEOPMENT"
}
func init() {
	fmt.Println(env, "env loaded.")
}
func main() {
	fmt.Println("main func after the inits")
}
