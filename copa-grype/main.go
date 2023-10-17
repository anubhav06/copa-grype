package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	grypeTypes "github.com/anchore/grype/grype/presenter/models"
	v1alpha1 "github.com/project-copacetic/copacetic/pkg/types/v1alpha1"
)

type GrypeParser struct{}

func parseGrypeReport(file string) (*grypeTypes.Document, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var gr grypeTypes.Document
	if err = json.Unmarshal(data, &gr); err != nil {
		return nil, err
	}

	return &gr, nil
}

func NewGrypeParser() *GrypeParser {
	return &GrypeParser{}
}

func (k *GrypeParser) Parse(file string) (*v1alpha1.UpdateManifest, error) {
	// Parse the grype scan results
	report, err := parseGrypeReport(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal function is not able to detect if the report is in the correct format. It returns no error even if the report is in the wrong format.
	// Therefore, we check if the report is in the correct format by parsing the Descriptor Name
	// If the name does not contain grype, then it is not in the correct format, report is marked as unsupported
	if report.Descriptor.Name != "grype" {
		return nil, errors.New("report format not supported by grype")
	}

	if err != nil {
		return nil, err
	}

	updates := v1alpha1.UpdateManifest{
		APIVersion: v1alpha1.APIVersion,
		Metadata: v1alpha1.Metadata{
			OS: v1alpha1.OS{
				Type:    report.Distro.Name,
				Version: report.Distro.Version,
			},
			Config: v1alpha1.Config{
				Arch: report.Source.Target.(map[string]interface{})["architecture"].(string),
			},
		},
	}

	// Check if vulnerability is OS-lvl package & check if vulnerability is fixable
	for i := range report.Matches {
		vuln := &report.Matches[i]
		if vuln.Artifact.Language == "" && vuln.Vulnerability.Fix.State == "fixed" {
			updates.Updates = append(updates.Updates, v1alpha1.UpdatePackage{Name: vuln.Artifact.Name, InstalledVersion: vuln.Artifact.Version, FixedVersion: vuln.Vulnerability.Fix.Versions[0], VulnerabilityID: vuln.Vulnerability.ID})
		}
	}
	return &updates, nil
}

func main() {
	// Initialize the parser
	grypeParser := NewGrypeParser()
	// Get the image report from command line
	imageReport := os.Args[1]

	report, err := grypeParser.Parse(imageReport)
	if err != nil {
		fmt.Printf("Error scanning image: %v\n", err)
		return
	}

	// Serialize the standardized report and print it to stdout
	reportBytes, err := json.Marshal(report)
	if err != nil {
		fmt.Printf("Error serializing report: %v\n", err)
		return
	}

	os.Stdout.Write(reportBytes)
}
