/*
- check and create folder
-
*/
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	mydir := "/tmp/abc/xyz/abc1"

	statmydir, err := os.Stat(mydir)

	// if no err then folder exists
	if err == nil {
		fmt.Println("folder exists", statmydir.Mode(), statmydir.IsDir())
		log.Fatal(mydir, " already exists")
	} else if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(mydir, 0777)
		if err != nil {
			log.Fatal("folder creation failed ", err)
		}
	}
	fmt.Println(mydir, " folder was created")

	// remove the folder to continue to next portion
	os.Remove(mydir)

	// folder listing
	// ReadDir(name string) ([]DirEntry, error)
	ls, lsErr := os.ReadDir("./")
	if lsErr != nil {
		log.Fatal("folder listing : ", lsErr)
	}
	fmt.Println("FILENAME,  DIR?")
	for _, f := range ls {
		fmt.Println(f.Name(), f.IsDir())
	}

}
