package packageone

import "fmt"

var PackageVar = "This is a package level variable in packageone."

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}


// Lower case first letter = only available in this package
// var privateVar = "I am private"

// // Upper case first letter = available any where in project
// var PublicVar = "I am public (or exported)"

// // This function is private
// func notExported() {

// }

// // This function is public
// func Exported() {

// }