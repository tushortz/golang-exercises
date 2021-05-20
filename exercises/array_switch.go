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
	// If the array contains 0 in both rows and
	// columns, change every row and column the
	// number falls on to 0

	// multidimentional slice declaration
	array := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1},
	}

	// slices to store the row and index position
	rows := []int{}
	cols := []int{}

	// get te rows and column index of the number 0 we are interested in
	for i, row := range array {
		for j, col := range row {
			if col == 0 {
				array[i][j] = 0
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	// do the replacement based on the rows and column index we want
	for i, row := range array {
		for j, _ := range row {
			if indexOf(i, rows) > -1 || indexOf(j, cols) > -1 {
				array[i][j] = 0
			}
		}
	}

	// print result
	fmt.Println(array)
}
