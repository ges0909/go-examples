package interface_tests

import "bytes"

type Letters interface {
	get() string
}

type A struct{}

func (o A) get() string {
	return "A"
}

type B struct{}

func (o B) get() string {
	return "B"
}

type C struct{}

func (o C) get() string {
	return "C"
}

func concat(letters []Letters) string {
	var buffer bytes.Buffer
	for _, l := range letters {
		buffer.WriteString(l.get())
	}
	return buffer.String()
}
