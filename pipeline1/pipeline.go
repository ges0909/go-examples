package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	// returns created channel immediately because writing to
	// channel is done in goroutine and therefore non-blocking
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	in := gen(2, 3, 4, 5, 6)
	out := sq(in)

	// for n, more := <-out; more; n, more = <-out {
	for n := range out {
		// Consume the output.
		fmt.Println(n)
	}
}
