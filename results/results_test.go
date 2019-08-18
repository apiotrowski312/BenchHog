package results

import (
	"reflect"
	"testing"
	"time"
)

func createValuesForTests(timeNeeded int, wasOk bool) chan Measurement {
	measurment := make(chan Measurement, 5)
	for i := 0; i < 5; i++ {
		measurment <- Measurement{waitTime: time.Duration(timeNeeded), success: wasOk}
	}

	close(measurment)
	return measurment
}

func Test_toSlice(t *testing.T) {
	type args struct {
		c chan Measurement
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []bool
	}{
		{
			"Basic Test",
			args{
				c: createValuesForTests(1, true),
			},
			[]int{1, 1, 1, 1, 1},
			[]bool{true, true, true, true, true},
		},
		{
			"Basic Test",
			args{
				c: createValuesForTests(2, false),
			},
			[]int{2, 2, 2, 2, 2},
			[]bool{false, false, false, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := toSlice(tt.args.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toSlice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("toSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
