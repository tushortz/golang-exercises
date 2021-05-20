package main

import "fmt"

func main() {
	i := 1

	for i <= 100 {
		if i%15 == 0 {
			fmt.Println(i, "Fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else {
			fmt.Println(i)
		}

		i += 1
	}
}
