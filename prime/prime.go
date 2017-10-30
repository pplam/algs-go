package main

import (
	"os"
	"strconv"
	"fmt"

	"./sieve"
)

func main() {
	s := os.Args[1]

	upper, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	primes := sieve.Sieve(upper)
	fmt.Println(primes)
}
