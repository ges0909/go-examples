package semaphore_with_chan

import (
	"log"
	"os"
	"time"
)

var Trace *log.Logger = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

const MaxOutstanding = 5

var semaphore = make(chan int, MaxOutstanding) // buffered channel

func Serve(no int, quit chan bool) {

	semaphore <- no                             // push to channel, wait when buffer is full
	Trace.Printf("-- len=%v\n", len(semaphore)) // number of elements queued in channel buffer

	// func literal is closure
	go func() {
		d := time.Duration(1) * time.Second // time.Duration holds nanoseconds
		// d := time.Duration(no) * time.Second
		time.Sleep(d)
		<-semaphore // pop from channel
		Trace.Printf("++ no=%v\n", no)
		quit <- true
	}()

}
