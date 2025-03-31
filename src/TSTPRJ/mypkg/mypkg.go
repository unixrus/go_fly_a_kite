/*
mypkg package
*/
package mypkg

// public funcs start with upperscase
func Hello(x string) (string, string) {
	return "(public) mypkg -> Hello ", x
}

// private funcs start with lowercase
func phello(x string) (string, string) {
	return "(private) mypkg -> phello ", x
}
