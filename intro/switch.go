package main

import "fmt"

func main() {
	switch val := 9; val > 9 {
	case true:
		fmt.Println("Short")
	case false:
		fmt.Println("Long")
	default:
		fmt.Println("Nothing")
	}
}
