/* 
TBD - untested
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func closethefile(f *os.File) error {
	f.Close()
	fmt.Println(f.Name(), " is closed")
	return nil
}

func main() {

	/*
		testfile := "/tmp/abc.txt"

			// create file.
			myfile, err := os.Create(testfile)
			defer closethefile(myfile)
			if err != nil {
				log.Fatal("normal file creation err: ", err)
			} else {
				fmt.Println(testfile, " created, pointer: ", myfile)
			}
			// force close the created file
			// else defer closes it
			closethefile(myfile)

			// remove file.
			err = os.Remove(testfile)
			if err != nil {
				log.Fatal(testfile, " failed to remove: ")
			} else {
				fmt.Println(testfile, " removed")
			}
	*/

	// open file
	testfile1 := "/tmp/abc1.txt"
	myfile1, err := os.Open(testfile1)
	// defer closethefile(myfile1)
	closethefile(myfile1) // not defer'ing as renaming it next
	if err != nil {
		log.Fatal("unable to ", err)
	} else {
		fmt.Println("file opened: ", myfile1.Name())
	}

	// rename file
	newtestfile1 := "/tmp/abc2.txt"

	err = os.Rename(testfile1, newtestfile1)
	if err != nil {
		log.Fatal("failed to rename: ", err)
	} else {
		fmt.Println("file renamed")
	}

}
