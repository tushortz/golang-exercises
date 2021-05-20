package main

import "fmt"

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func main() {
	array := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1},
	}

	rows := []int{}
	cols := []int{}

	for i, row := range array {
		for j, col := range row {
			if col == 0 {
				array[i][j] = 0
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	for i, row := range array {
		for j, _ := range row {
			if indexOf(i, rows) > -1 || indexOf(j, cols) > -1 {
				array[i][j] = 0
			}
		}
	}

	fmt.Println(array)
}
