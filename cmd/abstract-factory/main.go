package main

import (
	"design-patterns/creational"
	"fmt"
)

func main() {
	fmt.Println("-----------------------------")
	fmt.Println("Abstract Factory Pattern Demo")
	// Get report generator factory
	reportGeneratorFactory := creational.GetReportGeneratorFactory("pdf")
	if reportGeneratorFactory == nil {
		panic("Report generator factory is not defined.")
	}

	// Create report generator
	reportGenerator := reportGeneratorFactory.CreateReportGenerator()
	if reportGenerator == nil {
		panic("Report generator is not defined.")
	}
	reportGenerator.GenerateReport()
}
