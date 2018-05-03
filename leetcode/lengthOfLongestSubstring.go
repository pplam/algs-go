package main

import "fmt"

func main() {
	s := "pwwkew"
	l := lengthOfLongestSubstring(s)
	fmt.Println(l)
}

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)

	hit := func(p, h int) int {
		for q := p + 1; q < h; q++ {
			if runes[q] == runes[p] {
				return q
			}
		}
		return h
	}

	longest := 0
	start := 0
	stop := len(runes)
	for i := 0; i < len(runes); {
		for j := i; j < stop; j++ {
			pair := hit(j, stop)
			if pair < stop {
				stop = pair
				start = j
			}
		}

		length := stop - i
		if length > longest {
			longest = length
		}

		stop = len(runes)
		if start > i + 1 {
			i = start + 1
		} else {
			i++
		}
	}

	return longest
}
