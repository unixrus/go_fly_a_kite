/*
description
https://pkg.go.dev/strings
*/
package main

import (
	"fmt"
	"strings"
)

// https://pkg.go.dev/builtin package builtins

func main() {
	fmt.Println("Action ")
	fmt.Println(strings.Contains("hello this is a test 123A", "test"))
	fmt.Println(strings.Replace("test1 hello this is a test 123A", "test", "NoTest", 0))
	fmt.Println(strings.Replace("test1 hello this is a test 123A", "test", "NoTest", 1))
	fmt.Println(strings.Replace("test1 hello this is a test 123A", "test", "NoTest", -1))
	fmt.Println(strings.ReplaceAll("test1 hello this is a test 123A", "test", "NoTest"))
	fmt.Println(strings.ReplaceAll("test1 hello this is a test 123A", "test", ""))
	fmt.Println(strings.Title("test1 hello this is a test 123A"))
	fmt.Println(strings.ToUpper("test1 hello this is a test 123A"))
	fmt.Println(strings.Trim("_test1 hello this is a test 123A_", "_"))

}

/*
Action

*/
