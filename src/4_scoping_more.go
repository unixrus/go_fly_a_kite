/*
scoping rules ... vars dont change outside local scope
*/
package main

import "fmt"

// globall scope
var wrkday = 2

func main() {
	// local scope
	wrkday := 6

	switch wrkday {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 6:
		fmt.Println("Sturday")
	default:
		fmt.Println("holiday")
	} // end of switch
}

/*
local wrkday value is used in the switch
Saturdaay
*/
