package main

type T struct {
	i int
	s string
}

func getObj1() *T {
	p := new(T) // local scope
	p.i = 9961
	p.s = "ges0909"
	return p
}

func getObj2() *T {
	return &T{9961, "ges0909"} // local scope
}
