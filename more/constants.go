package main

import (
	"fmt"
)

func main() {
	const x = 45
	const (
		a = iota
		b = iota
		c = 13
		d = ""
		e = iota
	)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)

	const (
		dog = iota
		cat
		fish
	)

	fmt.Printf("%v, %T\n", cat == b, cat)

}
