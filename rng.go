package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func pseudoRandomGenerator(start, end, n int, delimiter string) (string, error) {
	if start > end {
		return "", fmt.Errorf("start should be less than or equal to end")
	}
	if n <= 0 {
		return "", fmt.Errorf("number of random numbers to generate should be greater than zero")
	}

	rngs := make([]int, n)
	for i := 0; i < n; i++ {
		rngs[i] = rand.Intn(end-start+1) + start
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(rngs)), delimiter), "[]"), nil
}
