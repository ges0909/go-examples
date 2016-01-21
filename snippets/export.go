// Demonstrates the visibility outside a package.
// Type- and member declaration starting with a
// capitalized letter are exported; otherwise
// not.

package main

// a is not visible outside package
type a struct {
	// 'n' not visible outside package
	n string
}

// B is visible outside package
type B struct {
	// 'N' visible outside package
	N string
}

// C is visible outside package
type C struct {
	// 'n' not visible outside package
	n string
}

// N getter
func (c *C) N() string {
	return c.n
}

// setN setter
func (c *C) setN(v string) {
	c.n = v
}
