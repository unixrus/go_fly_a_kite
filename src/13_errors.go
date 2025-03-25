/*
err handling
*/
package main

import (
	"errors"
	"fmt"
)

func themonth(numbr int) (string, error) {
	if numbr < 1 || numbr > 12 {
		// error as month go from 1 - 12 only
		return "None", errors.New("only 1-12 months")
	}
	switch numbr {
	case 1:
		return "Jan", nil
	case 2:
		return "Feb", nil
	case 3:
		return "Mar", nil
	case 4:
		return "Apr", nil
	case 5:
		return "May", nil
	case 6:
		return "Jun", nil
	case 7:
		return "Jul", nil
	case 8:
		return "Aug", nil
	case 9:
		return "Sep", nil
	case 10:
		return "Oct", nil
	case 11:
		return "Nov", nil
	case 12:
		return "Dec", nil
	default:
		// this will never happen due to if
		return "None", errors.New("only 1-12 months")
	}
}
func main() {

	mnth, err := themonth(13)
	if err == nil {
		fmt.Println(mnth)
	} else {
		fmt.Println(err)
	}
}
