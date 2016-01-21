package goroutine

import (
	"fmt"
	"time"
)

func sayWithFuncLiteral(msg string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(msg)
	}() // Note the parentheses - must call the function.
}

func say(msg string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(msg)
}

type chanDataType struct {
	msg   string
	delay time.Duration
}

func sayWithChannel(ch chan chanDataType) {
	d := <-ch
	time.Sleep(d.delay)
	fmt.Println(d.msg)
}
