package main

import "fmt"

func main() {
	days := map[string]int{
		"mon": 1,
		"tue": 2,
		"wed": 3,
	}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println("loop" + string(i))
		for k, v := range days {
			fmt.Println(i*v, k)
		}
	}

	str := []string{"hello", "world", ""}
	fmt.Println(str)

outer:
	for _, v := range str {
		fmt.Println(v)

		for i, r := range v {
			fmt.Println(i, r, string(r))
		}

		goto outer
	}
}
