package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)


func main() {
	file := os.Args[1]
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error in reading file")
		os.Exit(1)
	}
	mvalues := string(bytes)
	mvalues = strings.Trim(mvalues, "\n")
  data := strings.Split(mvalues, "\n")
  size:=len(data)
  fmt.Println(len(data))
	if !is_square_matrix(size) {
		fmt.Println("Invalid Matrix")
		os.Exit(1)
	}
	asize:= int(math.Sqrt(float64(size)))
	var matrix [10][10]int64

	for _, value := range data {
		lines := strings.Fields(value)
		row, _ := strconv.ParseInt(lines[0], 10, 64)
		col, _ := strconv.ParseInt(lines[1], 10, 64)
		val, _ := strconv.ParseInt(lines[2], 10, 64)
		matrix[row][col] = val
	}
	print_matrix(matrix, "A", asize)
	tmatrix := transpose_matrix(matrix, asize)
	print_matrix(tmatrix, "A-T", asize)
	result := multiply_matrices(matrix, tmatrix, asize)
	print_matrix(result, "A*A-T", asize)
}
func transpose_matrix(matrix [10][10]int64, size int) [10][10]int64 {
	var transpose [10][10]int64
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			transpose[i][j] = matrix[j][i]
		}
	}
	return transpose
}
func multiply_matrices(m, t [10][10]int64, size int) [10][10]int64 {
	var resMatrix [10][10]int64
	var resultValue int64
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				resultValue = resultValue + m[i][k]*t[k][j]
			}
			resMatrix[i][j] = resultValue
			resultValue = 0
		}

	}
	return resMatrix
}
func print_matrix(rmatrix [10][10]int64, name string, size int) {
	fmt.Println(name)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("\t", rmatrix[i][j], " ")
		}
		fmt.Println()
	}

}
func is_square_matrix(size int) bool {
	sqrt := int(math.Sqrt(float64(size)))
	return sqrt*sqrt == size
}
