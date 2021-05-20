package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := [...]int{1_000, 2, 3}
	y := []int{1_000, 2, 3}

	y[0] = 4
	fmt.Printf("%v %v %v\n", len(x[2:]), cap(x[2:]), x[2:])
	fmt.Printf("%v %v %v\n", len(y[2:]), cap(y[2:]), y[2:])

	fmt.Println(x[:])
	fmt.Println(y[:])

	z := make([]int, 3, 44)
	fmt.Printf("%v\n", z)
	fmt.Printf("%T\n", z)
	fmt.Printf("%v\n", len(z))
	fmt.Printf("%v\n", cap(z))

	w, e := strconv.Atoi("5")
	fmt.Println(w, e == nil)
}
