// channel
package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	n, more := <-ch
	for more {
		fmt.Println(n)
		n, more = <-ch
	}

	// similiar but more simple
	//	for n := range ch {
	//		fmt.Println(n)
	//	}

}
