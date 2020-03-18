package main

import "fmt"

type Stringer interface {
	String() string
}

type binary int

func Binary(i int) binary {
	var b = binary(i)
	return b
}

func (b binary) String() string {
	return fmt.Sprintf("%d", b)
}

func Conversion() {
	var b Stringer
	var i binary = 1
	b = i
	_ = b.String()
}

func Conversion2() {
	var b Stringer = Binary(1)
	_ = b.String()
}
