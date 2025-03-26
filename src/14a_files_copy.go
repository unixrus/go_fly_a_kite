/*
copy is not working TBD

error:
2025/03/26 16:56:57 write /tmp/src.txt: copy_file_range: bad file descriptor
exit status 1

https://github.com/golang/go/issues/60181 ?
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func closethefile(f *os.File) error {
	f.Close()
	fmt.Println(f.Name(), " is closed")
	return nil
}

func main() {

	// copy contents of src to dst file.

	srcf := "/tmp/src.txt"
	dstf := "/tmp/copy_of_5_slices.go"
	fmt.Println("copy ", srcf, " -> ", dstf)

	// open src file.
	// src, err := os.OpenFile(srcf, os.O_RDONLY, 0444)
	src, err := os.Open(srcf)
	// defer src.Close()
	defer closethefile(src)
	if err != nil {
		log.Fatal("src file", err)
	}

	// open/create destination file.
	dst, err := os.Create(dstf)
	closethefile(dst)
	// dst, err := os.OpenFile(dstf, os.O_WRONLY|os.O_CREATE, 0755)
	dst, err = os.OpenFile(dstf, os.O_WRONLY, 0777)
	// defer dst.Close()
	defer closethefile(dst)
	if err != nil {
		log.Fatal("dst file", err)
	}

	// copy the file contenst
	w, err := io.Copy(src, dst)
	if err != nil {
		fmt.Println("copy file err", err)
		// log.Fatal("copy file err", err)
	}
	// if w, err := io.Copy(src, dst); err != nil {
	// log.Fatal(err)
	// fmt.Println	fmt.Println(err)
	// }

	fmt.Println(reflect.TypeOf(w))
	fmt.Println(w)

}
