package main

import "fmt"

func main() {
	// array[start:end]
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	// append
	n = append(n, 30)
	fmt.Println(n)

	// 多次元配列
	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)
	fmt.Println(board[1])
	fmt.Println(board[1][2])
	m := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(m), cap(m), m)
}
