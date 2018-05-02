package main

import "fmt"

func main() {
	A := []int{1, 4, -1, 3, 2}
	length := solution(A, 5)
	fmt.Println(length)
}

func solution(A []int, N int) int {
    length := 1
		curr := 0
		next := A[curr]
		for next != -1 {
			curr = next
			next = A[curr]
			length++
    }
    return length
}
