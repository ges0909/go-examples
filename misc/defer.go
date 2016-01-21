package main

import (
	"errors"
	"fmt"
	"log"
)

func runtimeError() (err error) {
	// 1. A deferred func's arguments are evaluated when the defer statement is evaluated.
	// 2. Deferred func calls are executed in LIFO order after the surrounding function returns.
	// 3. Deferred func may read and assign to the returning function's named return values.
	defer func() {
		// Recover is only useful inside deferred func. During normal execution,
		// a call to recover will return nil and have no other effect.
		if e := recover(); e != nil { // check if deferred func is called because of panic
			switch t := e.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()

	tmp := 0
	result := 10 / tmp  // division by zero causes the desired panic
	fmt.Println(result) // avoid compiler error "result declared and not used"

	return nil
}

func main() {
	err := runtimeError()
	if err != nil {
		log.Fatal(err)
	}
}
