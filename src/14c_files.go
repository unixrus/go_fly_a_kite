/*
file create, rename, truncate and stat
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// open file
	testfile1 := "/tmp/abc1.txt"
	myfile1, err := os.Create(testfile1)
	if err != nil {
		log.Fatal("unable to ", err)
	} else {
		fmt.Println("file opened: ", myfile1.Name())
	}
	myfile1.Close()

	// rename file
	newtestfile1 := "/tmp/abc2.txt"

	err = os.Rename(testfile1, newtestfile1)
	if err != nil {
		log.Fatal("failed to rename: ", err)
	} else {
		fmt.Println("file renamed")
	}

	// truncate file to last 100bytes
	err = os.Truncate(newtestfile1, 100)
	if err != nil {
		log.Fatal("failed to trucate: ", err)
	} else {
		fmt.Println("file truncated")
	}

	// check file size
	f, err := os.Stat(newtestfile1)
	if err != nil {
		log.Fatal("failed to stat: ", err)
	} else {
		fmt.Println("size:", f.Size())
	}

}

/*
file opened:  /tmp/abc1.txt
file renamed
file truncated
size: 100
*/
