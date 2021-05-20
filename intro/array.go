package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var w [3]int
	var x [3]int = [3]int{10, 20, 30}
	var y = [3]int{10, 20, 30}
	var z = [12]int{1, 5: 4, 6, 10: 200, 15}

	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	r := true
	fmt.Printf("%v, %T", r, r)
}
