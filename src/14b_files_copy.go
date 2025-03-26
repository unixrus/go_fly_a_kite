/*
copy file contents
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// copy contents of src to dst file.

	srcf := "/tmp/src.txt"
	dstf := "/tmp/copy_of_5_slices.go"
	fmt.Println("copy ", srcf, " -> ", dstf)

	// read data from src file.
	// data, err := os.ReadFile(srcf)
	data, err := ioutil.ReadFile(srcf)
	if err != nil {
		log.Fatal("read src file", err)
	}

	// write destination file.
	// err = os.WriteFile(dstf, data, 0644)
	err = ioutil.WriteFile(dstf, data, 0644)
	if err != nil {
		log.Fatal("failed to write file", err)
	}
	fmt.Println("copy success")

}
