package results

import (
	"reflect"
	"testing"
	"time"
)

func createValuesForTests(timeNeeded int, status []int) chan Result {
	results := make(chan Result, len(status))
	for _, s := range status {
		results <- Result{requestTime: time.Duration(timeNeeded), responseStatus: s}
	}

	close(results)
	return results
}

func TestCreateResult(t *testing.T) {
	type args struct {
		requestTime    time.Duration
		responseStatus int
	}
	tests := []struct {
		name string
		args args
		want Result
	}{
		{
			"Basic test",
			args{time.Duration(0), 500},
			Result{requestTime: time.Duration(0), responseStatus: 500},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateResult(tt.args.requestTime, tt.args.responseStatus); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_GetResponseStatus(t *testing.T) {
	type fields struct {
		requestTime    time.Duration
		responseStatus int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"Basic get test",
			fields{requestTime: time.Duration(0), responseStatus: 500},
			500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Result{
				requestTime:    tt.fields.requestTime,
				responseStatus: tt.fields.responseStatus,
			}
			if got := m.GetResponseStatus(); got != tt.want {
				t.Errorf("Result.GetResponseStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateResultSummary(t *testing.T) {
	type args struct {
		c chan Result
	}
	tests := []struct {
		name string
		args args
		want ResultSummary
	}{
		{
			"Basic Test with 200 status code",
			args{
				c: createValuesForTests(1, []int{200, 200, 200, 200, 200}),
			},
			ResultSummary{
				timeStats: TimeStats{
					averageResponseTime: time.Duration(1),
					medianResponseTime:  time.Duration(1),
					responseTimePercantage: map[int]time.Duration{
						25:  time.Duration(1),
						50:  time.Duration(1),
						75:  time.Duration(1),
						100: time.Duration(1),
					},
				},
				statusCodes: StatusCodes{status2xx: 5, status3xx: 0, status4xx: 0, status5xx: 0},
			},
		},
		{
			"Test with 500",
			args{
				c: createValuesForTests(1, []int{200, 400, 500, 300, 500, 400, 500, 500, 500}),
			},
			ResultSummary{
				timeStats: TimeStats{
					averageResponseTime: time.Duration(1),
					medianResponseTime:  time.Duration(1),
					responseTimePercantage: map[int]time.Duration{
						25:  time.Duration(1),
						50:  time.Duration(1),
						75:  time.Duration(1),
						100: time.Duration(1),
					},
				},
				statusCodes: StatusCodes{status2xx: 1, status3xx: 1, status4xx: 2, status5xx: 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateResultSummary(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateResultSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
