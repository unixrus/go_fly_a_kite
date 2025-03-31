/*
file read and print

* determine type of var w/reflect.TypeOf
*/
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {

	testfile := "/tmp/abc.txt"
	// ReadFile(filename string) ([]byte, errors)
	fcontents, err := os.ReadFile(testfile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("--->>> HELLO: create temp file ", testfile)
		}
		log.Fatalln("err reading file: ", err)
	}

	// determine type of var
	fmt.Println("contents is type: ", reflect.TypeOf(fcontents))
	// contents is type:  []uint8

	fmt.Println(fcontents) // prints slice of bytes

	fmt.Println(string(fcontents)) // converts/typecast slice data to string
}

/*
 */
