package main

import "fmt"

// box is exportable obj
type Box struct {
	D float64
	W float64
	H float64
}

// volume method to calc vol of box
// methods includes a receiver.
// The receiver needs both a name and a type, in braces --> (b *Box)
// the name of receiver is b and the type is a pointer to a Box object.
// then the name of the method --> volume.
// method is going to return a float64 value,
// followed by curly braces
func (b *Box) volume() float64 {

	return b.D * b.W * b.H
}
func main() {

	// the box dimensions

	theBox := Box{D: 5, W: 7, H: 44}
	fmt.Println(theBox.volume())

}
