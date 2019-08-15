package main

import (
	"fmt"
	"sort"
	"time"
)

// PrintResults - print pretty and accurate metrics like mean time etc.
func PrintResults(results chan Measurement) {

	times, _ := toSlice(results)

	standardTimes(times)
	percentageTimes(times)
}

func standardTimes(times []int) {
	sort.Ints(times)

	total := 0
	for _, value := range times {
		total += value
	}

	average := time.Duration(float64(total) / float64(len(times)))
	median := time.Duration(times[len(times)/2])
	fmt.Println("Average time:\t", average)
	fmt.Println("Median time:\t", median)
}

func percentageTimes(times []int) {
	sort.Ints(times)
	twentyFive := time.Duration(times[len(times)/4])
	fifty := time.Duration(times[len(times)/2])
	seventyFive := time.Duration(times[len(times)/4*3])
	hundred := time.Duration(times[len(times)-1])

	fmt.Println()
	fmt.Println(" 25% time lower than:\t", twentyFive)
	fmt.Println(" 50% time lower than:\t", fifty)
	fmt.Println(" 75% time lower than:\t", seventyFive)
	fmt.Println("100% time lower than:\t", hundred)

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
