package deferring

import (
	"regexp"
	"testing"
)

func Test_1(t *testing.T) {
	err := runtimeError() // create a runtime divide by zero panic
	str := err.Error()
	exp := "runtime error: integer divide by zero"
	if str != exp {
		t.Error("Test failed")
	}
}

func Test_2(t *testing.T) {
	err := panicError() // create a runtime divide by zero panic
	str := err
	exp := "simulated error: integer divide by zero"
	if str != exp {
		t.Error("Test failed")
	}
}

func Test_3(t *testing.T) {
	str := returnValue1()
	exp := "deferred return"
	if str != exp {
		t.Error("Test failed")
	}
}

func Test_4(t *testing.T) {
	str := returnValue2()
	exp := "normal return"
	if str != exp {
		t.Error("Test failed")
	}
}

func Test_5(t *testing.T) {
	err := returnError()
	match, _ := regexp.MatchString(".*it didn't work$", err.Error())
	if !match {
		t.Error("Test failed")
	}
}
