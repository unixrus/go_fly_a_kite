/*
map literal data type
  - key value pairs akin dictionary in py
    map[ data type for key followed by value type]
    map [string]int

:=  is short assignment operator
*/
package main

// https://pkg.go.dev/builtin package builtins
import "fmt"

func main() {

	// zero map declared - string for key and int value
	var mymap map[string]int
	fmt.Println("zero map ", mymap)
	fmt.Println("zero map is nil? ", mymap == nil)
	printMyMap(mymap)
	// zero slice  []
	// zero slice is nil?  true

	// mymap["myvalue"] = 100   // fails since declared but not initialized
	// make non zero map with make()
	mymap1 := make(map[string]int) // make map of ints
	printMyMap(mymap)
	mymap1["item_value1"] = 2000
	fmt.Println("map ", mymap1)
	fmt.Println("this map is nil? ", mymap1 == nil)
	mymap = mymap1 // assign existing mymap1 to declared mymap map
	printMyMap(mymap)

	// declare and initialize using map literal
	mymap2 := map[string]int{
		"key1": 100,
		"key2": 200}
	fmt.Println("mymap2 ", mymap2)
	fmt.Println("this mymap2 is nil? ", mymap2 == nil)
	mymap2["third"] = 3000
	printMyMap(mymap2)

	element, isPresent := mymap2["key2"]
	fmt.Println("// key2 element", element, isPresent) // 200  true
	element1, isPresent := mymap2["key7"]
	fmt.Println("// key7 element", element1, isPresent) // 0 false

	printMyMap(mymap2)
	delete(mymap2, "key2") // delete key2 from mymap2
	printMyMap(mymap2)

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
