// Additional mypkg package - originally a nested pkg
// same as mypkg
package mypkg

// since we are merging nested pkg into mypkg
// dupl funcs need to be renamed
func Hellon(x string) (string, string) {
	return "(public) nstpkg -> Hellon ", x
}
