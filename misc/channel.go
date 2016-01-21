package main

import (
	"fmt"
	"time"
)

func writer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go writer(ch)
	go func() { // reader
		n, more := <-ch
		for more {
			fmt.Println(n)
			n, more = <-ch
		}
	}()
	d := time.Duration(1) * time.Second
	time.Sleep(d)
	fmt.Println("-- The End.")
}
