package main

import (
	"fmt"
	"sort"
	"time"
)

// PrintResults - print pretty and accurate metrics like mean time etc.
func PrintResults(results chan Measurement) {

	times, _ := toSlice(results)

	averageTime(times)
	medianTime(times)
}

func averageTime(times []int) {
	total := 0
	for _, value := range times {
		total += value
	}

	average := time.Duration(float64(total) / float64(len(times)))
	fmt.Println("Average time:", average)
}

func medianTime(times []int) {
	sort.Ints(times)
	median := time.Duration(times[len(times)/2])
	fmt.Println("Median time:", median)
}

// toSlice - converts Measurement channel to slices
func toSlice(c chan Measurement) ([]int, []bool) {
	times := make([]int, 0)
	success := make([]bool, 0)

	for i := range c {
		times = append(times, int(i.waitTime))
		success = append(success, i.success)
	}
	return times, success
}
