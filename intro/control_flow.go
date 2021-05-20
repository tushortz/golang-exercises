package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int = rand.Intn(10)
	fmt.Println(n)

	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number: ", n)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i < 100 {
		fmt.Println(i)
		i *= 2
	}

	for {
		fmt.Println(45)
		break
	}

	nums := []int{2, 4, 6, 8, 10}
	for _, v := range nums {
		fmt.Println(v)
	}
}
