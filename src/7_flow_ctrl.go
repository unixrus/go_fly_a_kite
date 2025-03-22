/*
flow control
  - if
  - for
  - switch
  - defer
  - break
  - continue
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

func main() {

	// executed after main function is returned
	defer fmt.Println("last thing to do in main")

	count := 0
	// init, condition and post
	// init can be outside the loop
	// post can be inside the loop
	for i := 0; i < 10; i++ {
		count += i
	}
	fmt.Println("count: ", count)

	count = 0
	// init, condition and post
	for i := 0; i < 10; i++ {
		if count > 20 {
			break // stop looping
			// continue to skip fwd.
		}
		count += i
	}
	fmt.Println("count: ", count)

	i := 0
outr_label:
	count = 0
	// init, condition and post
	for i < 10 {
		if count > 20 {
			i++
			goto outr_label
		}
		count += i // only incremented if ...
		i++
	}
	fmt.Println("count: ", count)

	if count < 5 {
		fmt.Println("less than 5")
	} else if count == 5 {
		fmt.Println("less than 5")
	} else {
		fmt.Println("more than 5")
	}

	// switch replacement of above if
	// only one case is processed.
	switch {
	case count < 5:
		fmt.Println("less than 5")
	case count == 5:
		fmt.Println("less than 5")
	default:
		fmt.Println("more than 5")
	}

	// another switch
	month := 2
	switch month {
	case 1:
		fmt.Println("Jan")
	case 2:
		fmt.Println("Feb")
	case 3:
		fmt.Println("Mar")
	default:
		fmt.Println("after winter")
	}
}

/*
Action

*/
