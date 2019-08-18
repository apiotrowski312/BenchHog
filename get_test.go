package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/apiotrowski312/benchHog/results"
)

func TestGet(t *testing.T) {

	var wg sync.WaitGroup
	type args struct {
		url        string
		limitRatio chan int
		wg         *sync.WaitGroup
		data       chan results.Measurement
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Basic Test",
			args{
				url:        "http://localhost:9696",
				limitRatio: make(chan int, 1),
				wg:         &wg,
				data:       make(chan results.Measurement, 1),
			},
		},
	}

	for _, tt := range tests {

		ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		l, _ := net.Listen("tcp", ":9696")
		ts.Listener = l
		ts.Start()
		defer ts.Close()

		t.Run(tt.name, func(t *testing.T) {
			tt.args.wg.Add(1)
			tt.args.limitRatio <- 1
			Get(tt.args.url, tt.args.limitRatio, tt.args.wg, tt.args.data)

			close(tt.args.limitRatio)
			close(tt.args.data)
		})
	}
}
