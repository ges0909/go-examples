package main

import "fmt"

type MyClass struct {
	name string
}

func (object MyClass) printMemberValue() {
	fmt.Println("class: " + object.name)
}
