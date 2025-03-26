/*
error handling
- use errors pkg
*/
package main

import (
	"errors"
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

	// use errors pkg for known conditions
	if _, err := os.Open("/tmp/abc1.txt"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("custom err msg --> file does not exist")
		} else {
			log.Println(err) // other than a missing file error
		}
		return
	}
	fmt.Print("file opened successfully\n")
}
