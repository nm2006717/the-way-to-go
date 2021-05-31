package main

import "fmt"

func main() {

	mat := [][]int{{1, 2}, {2, 2}}
	isToe:= isToeplitzMatrix(mat)
	fmt.Println(isToe)

}

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	mm := matrix[0]
	for i := 1; i <= m-1; {
		for j := 0; j < n-1; {
			if mm[j] != matrix[i][j+1] {
				return false
			}
			j++
		}
		mm = matrix[i]
		i++
	}

	return true
}
