package results

import (
	"fmt"
	"sort"
	"time"
)

// Result - struct created to agregate all results
type Result struct {
	requestTime    time.Duration
	responseStatus int
}

// ResultSummary - struct for printing results
type ResultSummary struct {
	statusCodes StatusCodes
	timeStats   TimeStats
}

// StatusCodes - part of ResultSummary struct. It keep informations with number of status codes.
type StatusCodes struct {
	status2xx int
	status3xx int
	status4xx int
	status5xx int
}

// TimeStats - part of ResultSummary struct. It keep informations about timings.
type TimeStats struct {
	averageResponseTime    time.Duration
	medianResponseTime     time.Duration
	responseTimePercantage map[int]time.Duration
}

// CreateResult - creates struct
func CreateResult(requestTime time.Duration, responseStatus int) Result {
	return Result{requestTime, responseStatus}
}

// GetResponseStatus - return success from Results
func (m Result) GetResponseStatus() int {
	return m.responseStatus
}

// CreateResultSummary - creates struct
func CreateResultSummary(c chan Result) ResultSummary {
	times := make([]int, 0)
	timeSum := 0
	status2xx := 0
	status3xx := 0
	status4xx := 0
	status5xx := 0

	for i := range c {
		times = append(times, int(i.requestTime))
		timeSum += int(i.requestTime)
		if i.responseStatus < 300 {
			status2xx++
		} else if i.responseStatus < 400 {
			status3xx++
		} else if i.responseStatus < 500 {
			status4xx++
		} else {
			status5xx++
		}
	}
	statusCodes := StatusCodes{status2xx: status2xx, status3xx: status3xx, status4xx: status4xx, status5xx: status5xx}

	sort.Ints(times)

	timeLength := len(times)

	twentyFive := time.Duration(times[timeLength/4])
	fifty := time.Duration(times[timeLength/2])
	seventyFive := time.Duration(times[timeLength/4*3])
	hundred := time.Duration(times[timeLength-1])

	responseTimePercantage := map[int]time.Duration{
		25:  twentyFive,
		50:  fifty,
		75:  seventyFive,
		100: hundred,
	}
	average := time.Duration(float64(timeSum) / float64(timeLength))
	median := time.Duration(times[timeLength/2])

	timeStats := TimeStats{averageResponseTime: average, medianResponseTime: median, responseTimePercantage: responseTimePercantage}

	return ResultSummary{statusCodes, timeStats}
}

// PrintResults - print pretty and accurate metrics like: mean time etc.
func PrintResults(results chan Result) {
	summary := CreateResultSummary(results)

	timeStats := summary.timeStats
	statusCodes := summary.statusCodes
	fmt.Printf("\n")
	fmt.Println("Average time:\t", timeStats.averageResponseTime)
	fmt.Println("Median time:\t", timeStats.medianResponseTime)
	fmt.Println()
	fmt.Println("Response times:")
	for k, v := range timeStats.responseTimePercantage {
		fmt.Printf("%d%% time lower than:\t%v\n", k, v)
	}
	fmt.Println()
	fmt.Printf("Requests with status 2xx:\t%d\n", statusCodes.status2xx)
	fmt.Printf("Requests with status 3xx:\t%d\n", statusCodes.status3xx)
	fmt.Printf("Requests with status 4xx:\t%d\n", statusCodes.status4xx)
	fmt.Printf("Requests with status 5xx:\t%d\n", statusCodes.status5xx)
	fmt.Println()
}
