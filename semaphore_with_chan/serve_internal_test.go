package semaphore_with_chan

import (
	"testing"
)

func Test_1(t *testing.T) {

	quit := make(chan bool)

	// start goroutines
	for n := 1; n <= 2*MaxOutstanding; n++ {
		Serve(n, quit)
	}

	// wait on end of all goroutines
	var i = 0
	for <-quit {
		i++
		if i == 2*MaxOutstanding {
			break
		}
	}

}
