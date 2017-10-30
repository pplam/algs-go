package sieve

import (
	"sort"
)

func Sieve(upper int) []int {
	if upper < 2 {
		return []int{}
	}

	result := []int{2}

	low := 2
	high := 2 * 2
	for upper > low {
		if upper < high {
			high = upper
		}

		ch := make(chan int)
		defer close(ch)

		for x := low + 1; x <= high; x++ {
			go test(x, result, ch)
		}

		primes := []int{}
		counter := 0
		total := high - low
		for {
			if counter == total {
				break
			}

			prime := <-ch
			if prime > 0 {
				primes = append(primes, prime)
			}
			counter++
		}

		sort.Ints(primes)
		result = append(result, primes...)

		low = high
		high = low * low
	}

	return result
}

func test(x int, known []int, ch chan<- int) {
	for _, p := range known {
		if p * p > x {
			ch <- x
			return
		}
		if x % p == 0 {
			ch <- 0
			return
		}
	}
	ch <- x
}
