package main

import "design-patterns/creational"

func main() {
	logger := creational.GetLoggerInstance()

	// Add logs
	logger.AddLog("This is log entry 1.")
	logger.AddLog("This is log entry 2.")

	// Print logs
	logger2 := creational.GetLoggerInstance()

	logger2.ShowLogs()
}
