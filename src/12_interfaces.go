/*
interfaces  TBD
*/
package main

import "fmt"

type Autos interface {
	Fuel() string
	Size() string
}

// Prius is exportable obj
type Prius struct {
}

func (a Prius) Fuel() string {
	return "Prius is Hybrid fue efficient!"
}

func (a Prius) Size() string {
	return "Prius is a compact car!"
}

// Suburban is exportable obj
type Suburban struct {
}

func (c Suburban) Fuel() string {
	return "Suburban is gas guzzler!"
}

func (c Suburban) Size() string {
	return "Suburban is a large SUV!"
}

func main() {
	automobiles := []Autos{Prius{}, Suburban{}}
	for _, a := range automobiles {
		fmt.Println(a.Fuel(), a.Size())
	}
	/*
	Prius is Hybrid fue efficient! Prius is a compact car!
	Suburban is gas guzzler! Suburban is a large SUV!
	*/

}
