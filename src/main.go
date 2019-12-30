package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/apiotrowski312/benchHog/src/helpers"
)

// CommandDetails - TODO
type CommandDetails struct {
	ratio         int
	numOfRequests int
	jsonString    string
	jsonFile      string
	links         []string
}

func main() {

	// Flags
	commandDetails := CommandDetails{}

	flag.IntVar(&commandDetails.numOfRequests, "r", 100, "Provide number of requests")
	flag.StringVar(&commandDetails.jsonString, "body", "", "Json for post request")
	flag.StringVar(&commandDetails.jsonFile, "bodyFile", "", "File (with json usually) for post request (higher priority than string)")
	flag.IntVar(&commandDetails.ratio, "ratio", 10, "Provide ratio")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Provide method and links")
		os.Exit(1)
	}

	links := make([]string, len(flag.Args())-1)
	for index, link := range flag.Args()[1:] {
		links[index] = helpers.ParseLink(link)
	}

	commandDetails.links = links

	switch flag.Args()[0] {
	case "help":
		helpCommand()
	case "get":
		getBenchmarkCommand(commandDetails)
	case "post":
		postBenchmarkCommand(commandDetails)
	default:
		helpCommand()
	}
}

func helpCommand() {
	fmt.Printf("Available commands:\n")
	fmt.Printf("help\tshow help\n")
	fmt.Printf("get\t start benchmark with get request\n")   // TODO: descriptions
	fmt.Printf("post\t start benchmark with post request\n") // TODO: descriptions
}
