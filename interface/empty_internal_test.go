package interface_tests

import "testing"

func Test_1(t *testing.T) {
	letters := []Letters{A{}, new(B), C{}}
	exp := "ABC"
	if str := concat(letters); str != exp {
		t.Error("Test failed")
	}
}
