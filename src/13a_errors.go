/*
error handling
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// signature func Open(name string) (file *file, err error)
	thefile, err := os.Open("/tmp/abc.txt")

	// matching a close above opened file
	defer thefile.Close()

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("file opened:", thefile.Name())

}
