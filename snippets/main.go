package main

// Logging see http://www.goinggo.net/2013/11/using-log-package-in-go.html.
import (
	"fmt"
	"time"
)

func main() {

	//	// export.go
	//	a := a{n: "a"}
	//	fmt.Println("1." + a.n)

	//	b := B{"B"}
	//	fmt.Println("2." + b.N)

	//	c := C{}
	//	c.setN("cC")
	//	fmt.Println("3." + c.N())

	//	// scope.go
	//	fmt.Println(getObj1())
	//	fmt.Println(getObj2())

	//	// new.go
	//	var a t
	//	a.x = 1
	//	a.y = "abc"

	//	b := t{2, "ABC"}
	//	fmt.Printf("%v\n", b)

	//	var c *t = new(t)                // allocate zeroed memory
	//	fmt.Printf("%v, %v\n", c.x, c.y) // zeroed
	//	c.x = 4                          // c->x = 4
	//	c.y = "DEF"                      // c->y = "DEF"

	//	var d *t = &b
	//	fmt.Println(d)

	//	m := val()
	//	fmt.Println(*m)

	// goroutine.go
	say("Wiedersehen!", 2000*time.Millisecond)
	fmt.Print("Auf ")
	time.Sleep(4000 * time.Millisecond)

}
