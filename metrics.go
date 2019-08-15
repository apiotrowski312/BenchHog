package main

import (
	"fmt"
	"time"
)

// Metrics - print pretty and accurate metrics like mean time etc.
func Metrics(waitTime chan time.Duration) {
	results := toSlice(waitTime)

	averageTime(results)
}

func averageTime(results []time.Duration) {
	total := 0.0
	for _, value := range results {
		total += float64(value)
	}

	average := time.Duration(total / float64(len(results)))
	fmt.Println("Average time:", average)
}

// toSlice - converts time.Duration channel to slice
func toSlice(c chan time.Duration) []time.Duration {
	s := make([]time.Duration, 0)
	for i := range c {
		s = append(s, i)
	}
	return s
}
