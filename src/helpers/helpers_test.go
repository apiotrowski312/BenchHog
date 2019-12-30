package helpers

import (
	"io/ioutil"
	"os"
	"testing"
)

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
			if got := ParseLink(tt.args.link); got != tt.want {
				t.Errorf("parseLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_showLoader(t *testing.T) {
	type args struct {
		currentNumber int
		maxNumber     int
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			"Basic test",
			args{1, 4},
			"\033[G\t[##________]\t25%",
		},
		{
			"Basic test",
			args{3, 4},
			"\033[G\t[#######___]\t75%",
		},
		{
			"Basic test",
			args{13, 97},
			"\033[G\t[#_________]\t13%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			ShowLoader(tt.args.currentNumber, tt.args.maxNumber)

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout

			if string(out) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, out)
			}
		})
	}
}
