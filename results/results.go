package results

import (
	"fmt"
	"sort"
	"time"
)

// Measurement - struct created to agregate all measurments
type Measurement struct {
	waitTime time.Duration
	success  bool
}

func CreateMeasurment(waitTime time.Duration, success bool) Measurement {
	return Measurement{waitTime: waitTime, success: success}
}

// PrintResults - print pretty and accurate metrics like mean time etc.
func PrintResults(results chan Measurement) {
	times, successes := toSlice(results)

	standardTimes(times)
	percentageTimes(times)
	percentageSuccess(successes)
	fmt.Println()

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

func percentageSuccess(success []bool) {
	total := 0
	for _, wasOk := range success {
		if wasOk {
			total++
		}
	}

	fmt.Println()
	fmt.Printf("Requests success:\t%d\n", total)
	fmt.Printf("Requests failed:\t%d\n", len(success)-total)
	fmt.Println()
	fmt.Printf("%.2f%% of requests responsed with 2xx status code\n", float64(total)/float64(len(success))*100)
}

func toSlice(c chan Measurement) ([]int, []bool) {
	times := make([]int, 0)
	success := make([]bool, 0)

	for i := range c {
		times = append(times, int(i.waitTime))
		success = append(success, i.success)
	}
	return times, success
}
