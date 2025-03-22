/*
	pointers to pass references to values and records within code

:=  is short assignment operator
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

func main() {

	// pointer initialized
	var myptr *int
	fmt.Println("//declared ptr ", myptr)
	fmt.Println("myptr is nil? ", myptr == nil)
	// pointer initialized
	var somevar = 100 //  to assign int
	var myptr1 *int = &somevar
	fmt.Println("//declared ptr ", myptr1)
	fmt.Println("myptr is nil? ", myptr1 == nil)
	// myptr mem loc and data at the location
	myptr = myptr1
	fmt.Println("myptr  addr and data", myptr, *myptr)

	mystr := "my_string"
	var mystrptr *string // string pointer
	mystrptr = &mystr    // assign mystr location
	// above 2 lines in single decl and assignment
	// var mystrptr *string = &mystr
	fmt.Println("string and location of var:", mystr, &mystr)
	fmt.Println("location & data:", mystrptr, *mystrptr)
	// string and location of var: my_string 0xc000012090
	// location & data: 0xc000012090 my_string

	var ptr *int = new(int)
	fmt.Println("ptr and *ptr", ptr, *ptr)
	*ptr = 1010 // change data at pointer location
	fmt.Println("ptr and *ptr", ptr, *ptr)
	// ptr and *ptr 0xc000098058 0
	// ptr and *ptr 0xc000098058 1010

	somevar1 := 1011
	ptr_to_somevar1 := &somevar1
	// OR like this --> var ptr_to_somevar1 = &somevar1
	fmt.Println("mem location", ptr_to_somevar1)
	fmt.Println("data at location", *ptr_to_somevar1)
	// mem location 0xc000098068
	// data at location 1011
}

func printMyMap(themap map[string]int) {
	fmt.Printf("// len=%d %v\n", len(themap), themap)
}

/*
$  go run 5_maps.go
zero map  map[]
zero map is nil?  true
// len=0 map[]
// len=0 map[]
map  map[item_value1:2000]
this map is nil?  false
// len=1 map[item_value1:2000]
mymap2  map[key1:100 key2:200]
this mymap2 is nil?  false
// len=3 map[key1:100 key2:200 third:3000]
// key2 element 200 true
// key7 element 0 false
// len=3 map[key1:100 key2:200 third:3000]
// len=2 map[key1:100 third:3000]

more experiments --> https://go.dev/tour/moretypes/15

*/
