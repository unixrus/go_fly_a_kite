package main

// separate std package and local pkgs with a blank line
import (
	"fmt"

	"TSTPRJ1/mypkg"
)

func main() {

	mystr := "world!"

	// accessor is last part of pkg ..
	// i.e. mypkg and nstpkg

	// use mypkg Hello
	fmt.Println(mypkg.Hello(mystr))
	// (public) mypkg -> Hello  world!

	// use mypkg/nstpkg Hello func
	fmt.Println(mypkg.Hellon(mystr))
	// (public) nstpkg -> Hello  world!

}
