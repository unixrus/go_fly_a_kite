// test project of creating a pkgs

/*

new use of doc.go ...

$ pwd
TSTPRJ1

merged nested pkg nstpkg into mypkg

renamed duplicate func Hello to Hellon

$ go run main.go
(public) mypkg -> Hello  world!
(public) nstpkg -> Hellon  world!

// reading pkg documentation

$ go doc mypkg
package mypkg // import "TSTPRJ1/mypkg"

mypkg package - sample test package to invoke and test

Additional mypkg package - originally a nested pkg same as mypkg

func Hello(x string) (string, string)
func Hellon(x string) (string, string)

*/
