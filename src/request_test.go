package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/apiotrowski312/benchHog/src/results"
)

func TestGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want results.Result
	}{
		{
			"Should succeed",
			args{
				url: "http://localhost:9696",
			},
			results.CreateResult(
				time.Duration(0),
				200,
			),
		},
		{
			"Should failed (wrong port)",
			args{
				url: "http://localhost:1234",
			},
			results.CreateResult(
				time.Duration(0),
				500,
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
			if got := Get(tt.args.url); !reflect.DeepEqual(got.GetResponseStatus(), tt.want.GetResponseStatus()) {
				t.Errorf("Get() = %v, want %v", got.GetResponseStatus(), tt.want.GetResponseStatus())
			}
		})
	}
}
