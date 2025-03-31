/*
pkg imports
blank import i.e. string in line 14
named imports - line 16
using default accessor - line 30
dot imports line 17
*/
package main

import (
	"fmt"
	"os"
	_ "strings"

	"TSTPRJ1/mypkg"
	mynew "TSTPRJ1/mypkg"
	// dot import not recommended
	// . "os"
	// then simply use Stdout instead of os.Stdout
)

func main() {
	fmt.Println("blank imported but unused pkg -> strings")
	fmt.Println(mypkg.Hello("TEST ORIGINAL"))

	// could be a new or local pkg mynew
	fmt.Println(mynew.Hello("TEST mynew alias"))

	// use Fprintln instead
	fmt.Fprintln(os.Stdout, "Alt way of fmt.Println")
}

/*

$ go run 16e_pkgs.go
blank imported but unused pkg -> strings
(public) mypkg -> Hello  TEST ORIGINAL
(public) mypkg -> Hello  TEST mynew alias
Alt way of Println
*/
