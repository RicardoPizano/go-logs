package main

import "github.com/RicardoPizano/go-logs/logger"

func main() {
	logger.CreateLogger(logger.Configuration{
		EnableConsole: true,
		ConsoleLevel:  "1",
		EnableFile:    true,
		FileLevel:     "1",
		FileLocation:  "example/example_log",
	})
	logger.Info("main", "main", "log example")
}
