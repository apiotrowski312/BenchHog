package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/apiotrowski312/benchHog/results"
)

func TestGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want results.Measurement
	}{
		{
			"Should succeed",
			args{
				url: "http://localhost:9696",
			},
			results.CreateMeasurment(
				time.Duration(0),
				true,
			),
		},
		{
			"Should failed (wrong port)",
			args{
				url: "http://localhost:1234",
			},
			results.CreateMeasurment(
				time.Duration(0),
				false,
			),
		},
	}

	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	l, _ := net.Listen("tcp", ":9696")
	ts.Listener = l
	ts.Start()
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.url); !reflect.DeepEqual(got.GetSuccess(), tt.want.GetSuccess()) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
