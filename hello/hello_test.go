// package main
package hello_test

// see tips and tricks fpr go testing
// https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742#.bugstfj9o

import "testing" // test support
import "hello"   // package under test

func TestHello_CountArgsAsFunc(t *testing.T) { // the test gets the name of the package
	// test 1
	args := hello.MyType{"prog", "opts", "args"}
	actual, expected := hello.CountArgs(args), 3
	if actual != expected {
		t.Error("Test failed")
	}
}

func TestHello_CountArgsAsMethod(t *testing.T) { // the test gets the name of the package
	//test 2
	var obj hello.MyType
	obj = hello.MyType{"prog"}
	actual, expected := obj.CountArgs(), 1
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, actual)
	}
}
