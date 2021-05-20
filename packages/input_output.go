package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("packages/numbers.txt")
	b1 := make([]byte, 11)

	f.Read(b1)

	for i := range b1 {
		i := int(i)
		fmt.Println(i * i)
	}

	fmt.Println()

	f2, _ := ioutil.ReadFile("packages/numbers.txt")
	fmt.Println(string(f2))

}
