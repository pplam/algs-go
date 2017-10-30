package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello")
}

func filter(stop int) {
	ret := []int{2}
	fmt.Println(ret)
}

func test(x int, known []int) bool {
	for _, p := range known {
		if p * p > x {
			return true
		}
		if x % p == 0 {
			return false
		}
	}
	return true
}
