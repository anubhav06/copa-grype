package main

import (
	"encoding/json"
	"fmt"
	"os"

	copaGrype "github.com/anubhav06/copa-grype/grype"
)

func main() {
	if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <image_report>\n", os.Args[0])
        os.Exit(1)
    }

	// Initialize the parser
	grypeParser := copaGrype.NewGrypeParser()
	// Get the image report from command line
	imageReport := os.Args[1]

	report, err := grypeParser.Parse(imageReport)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error scanning image: %v\n", err)
		return
	}

	// Serialize the standardized report and print it to stdout
	reportBytes, err := json.Marshal(report)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error serializing report: %v\n", err)
		return
	}

	os.Stdout.Write(reportBytes)
}
