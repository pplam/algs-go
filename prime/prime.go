package main

import (
	"os"
	"strconv"
	"fmt"

	"./filter"
)

func main() {
	s := os.Args[1]

	stop, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	primes := filter.Filter(stop)
	fmt.Println(primes)
}
