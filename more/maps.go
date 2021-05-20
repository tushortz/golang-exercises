package main

import "fmt"

func main() {
	populations := map[string]int{
		"California": 1235,
		"Texas":      653,
		"Florida":    78953,
	}

	jack := populations
	item, ok := populations["t"]
	delete(jack, "Texas")

	fmt.Println(ok, item)
	fmt.Println(populations)
	fmt.Println(jack)
}
