// anonymous
package main

import (
	"fmt"
)

func main() {

	f := func(x int) int {
		return x * x
	}

	fmt.Println(f(9))
}
