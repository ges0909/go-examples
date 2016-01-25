package main

import (
	"fmt"
	"time"
)

func writer() {
	fmt.Println("function")
}

func main() {
	go writer() // function
	go func() { // closure
		fmt.Println("closure")
	}()
	d := time.Duration(1) * time.Second
	time.Sleep(d)
	fmt.Println("-- The End.")
}
