package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func Test_1(t *testing.T) {
	sayWithFuncLiteral("Wiedersehen!", 2000*time.Millisecond)
	fmt.Print("Auf ")
	time.Sleep(4000 * time.Millisecond)
}

func Test_2(t *testing.T) {
	go say("Wiedersehen!", 2000*time.Millisecond)
	fmt.Print("Auf ")
	time.Sleep(4000 * time.Millisecond)
}

func Test_3(t *testing.T) {
	d := chanDataType{"Hallo Welt!", 2000 * time.Millisecond}
	ch := make(chan chanDataType)
	go sayWithChannel(ch)
	ch <- d
	time.Sleep(4000 * time.Millisecond)
}
