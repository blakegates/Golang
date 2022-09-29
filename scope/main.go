package main

import (
	"myapp/packageone"
)

var myVar = "This is a package level variable."

func main() {
	var blockVar = "This is a block level variable."

	packageone.PrintMe(myVar, blockVar)
}

// variables
// declare a package level variable for the main package named myVar -check

// declare a block level variable for the main function called blockVar -check

// declare a package level variable in the packageone package named PackageVar -check

// create an exported function in packageone named PrintMe -check

// in the main function, print out the values of myVar, blockVar and PackageVar
// on one line using the PrintMe function in packageone -check
