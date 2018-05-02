package main

import "fmt"

func main() {
	A := []int{1, 2, 4, 7, 6, 6, 6, 2, 0}
	num := solution(A, 9)
	fmt.Println(num)
}

func solution(A []int, N int) int {
	if N == 0 || N == 1 {
		return 0
	}
	if N == 2 {
		return 1
	}

	rel0 := cmp(A[0], A[1])
	succ := 2
	count := 1
	for i := 2; i < N; i++ {
		rel := cmp(A[i - 1], A[i])
		if rel == rel0 {
			count += succ
			succ++
		} else {
			count++
			rel0 = rel
			succ = 2
		}
	}

	return count
}

func cmp(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
