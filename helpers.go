package main

import "strings"

func parseLink(link string) string {
	if !strings.HasPrefix(link, "http") {
		return "http://" + link
	}
	return link
}
