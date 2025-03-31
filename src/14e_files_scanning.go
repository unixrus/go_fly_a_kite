/*
file read line-by-line

* determine type of var w/reflect.TypeOf
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	testfile := "/tmp/abc.txt"

	// create a reader obj
	readf, err := os.Open(testfile)
	defer readf.Close()

	if err != nil {
		log.Fatal(err) // if issues with opening file.
	}

	// buffio implements bufferring I/O using reader & scanner obj types
	// bufio.Reader implements buffering for a Reader object,
	// and Scanner implements an interface for reading data such as lines of text in a file
	// with line breaks etc to to indicate EOL.

	// NewScanner func is passed a Reader object and it returns a Scanner
	// func newscanner (r io.Reader) *Scanner
	s := bufio.NewScanner(readf)

	// for loop to scan thru tokens till EOF
	for s.Scan() {
		// text method on scan token
		fmt.Println(s.Text())
	}

	// check for errors with scan
	err1 := s.Err()
	if err1 != nil {
		log.Fatal("err during scan: ", err)
	}
}

/*
 */
