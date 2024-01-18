package main

import "fmt"

/*
118. Pascal's Triangle

[Easy]

Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown: */

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = 1
	if numRows > 1 {
		dp[1][0] = 1
		dp[1][1] = 1
	}
	for i := 2; i < numRows; i++ {
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp
}

func main() {
	fmt.Printf("%v", generate(10))
}
