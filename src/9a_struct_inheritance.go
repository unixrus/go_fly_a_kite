/*
methods, structs, inheritance - needs more TLC
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

type Part struct {
	Manufacturer string
}

func (p *Part) Mfc() string {
	return p.Manufacturer
}

type Tyre struct {
	Part // anonymous field ?? TBD
}

func main() {
	fmt.Println("Action ")
	t := Tyre{Part{Manufacturer: "michelin"}}
	fmt.Println("manufacturuer of tire: ", t.Mfc())
}

/*
Action

*/
