package main

import "testing"

func Test_parseLink(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Parse link with http",
			args{
				link: "http://wp.pl",
			},
			"http://wp.pl",
		},
		{
			"Parse link witout http",
			args{
				link: "wp.pl",
			},
			"http://wp.pl",
		},
		{
			"Parse link wit https",
			args{
				link: "https://wp.pl",
			},
			"https://wp.pl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLink(tt.args.link); got != tt.want {
				t.Errorf("parseLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
