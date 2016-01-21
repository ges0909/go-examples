package deferring

import (
	"errors"
	"fmt"
	"time"
)

/*
 * defer + runtime error
 */
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

/*
 * defer + panic
 */
func panicError() (err interface{}) {

	defer func() {
		if e := recover(); e != nil { // check if panic arised; otherwise normal call of deferred func att end of sourrounding func
			err = e
		}
	}()

	panic("simulated error: integer divide by zero")

	return nil
}

/*
 * defer + return value
 */
func returnValue1() (str string) {
	str = "normal return"
	defer func() {
		str = "deferred return"
	}()
	return
}

/*
 * defer + return value
 */
func returnValue2() (str string) {
	str = "normal return"
	defer func() {
		if e := recover(); e != nil {
			str = "deferred return"
		}
	}()
	return
}

/*
 * define an own error
 */
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func returnError() error {
	return &MyError{time.Now(), "it didn't work"}
}
