package helpers

import (
	"fmt"
	"strings"
)

func ParseLink(link string) string {
	if !strings.HasPrefix(link, "http") {
		return "http://" + link
	}
	return link
}

func ShowLoader(currentNumber int, maxNumber int) {
	percentage := int(float64(currentNumber) / float64(maxNumber) * 100)

	hashes := strings.Repeat("#", percentage/10)
	underscore := strings.Repeat("_", 10-(percentage/10))

	fmt.Printf("\033[G\t[%s%s]\t%d%%", hashes, underscore, percentage)
}
