package main

type t struct {
	x int
	y string
}

func val() *int {
	n := 999
	return &n // not freed when leaving scope
}
