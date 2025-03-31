/*
blank import of string in line 8
*/
package main

import (
	"fmt"
	_ "strings"
)

func main() {
	fmt.Println("blank imported but unused pkg -> strings")
}
