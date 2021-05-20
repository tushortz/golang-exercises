package main

import (
	"fmt"
)

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x, y := 5, 5
		fmt.Println(x)
		fmt.Println(y)

	}
	fmt.Println(x)
	fmt.Println(true)
	true := 5
	fmt.Println(true)
}
