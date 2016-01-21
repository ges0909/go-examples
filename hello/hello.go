// First steps in go programming.

// To document a type, variable, constant, function, or even a package, write
// a regular comment directly preceding its declaration, with no intervening
// blank line. Godoc will then present that comment as text alongside the item
// it documents.

// This is a package declaration.
// package main
package hello

// These are package imports. Unused imports a flagged as error by the compiler.
import (
	"fmt"
	"os"
)

// Type declaration.
type MyType []string

// Method bounded to the  type definition. Type and method definition forming
// a class definition as known from traditional OO.
func (args MyType) CountArgs() int {
	return len(args)
}

// Simple function definition. Note the function has the same name like the
// method above.
func CountArgs(args MyType) int {
	return len(args)
}

// The main function of the main package is the program entry point. Unlike
// as in C/C++ or Java it has no arguments and return value.
func main( /*no arguments*/ ) /*no return value*/ {

	// Command line arguments are provided by os.Args as []string.
	if len(os.Args) == 1 {
		fmt.Println("Hallo Welt!")
	} else {
		for i := 1; i < len(os.Args); i++ {
			fmt.Print(os.Args[i])
			if i+1 < len(os.Args) {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}

	// This is a method call.
	var args MyType
	args = os.Args
	n1 := args.CountArgs()
	fmt.Printf("n1 = %v\n", n1) // %v - value in default format

	// This is a function call.
	n2 := CountArgs(os.Args)
	fmt.Printf("n2 = %v\n", n2)
}
