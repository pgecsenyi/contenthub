package main

import "flag"
import (
	"configuration"
	"server"
)

func main() {

	config := parseCommandLineArguments()

	server := server.Server{}
	server.Initialize(config)
	server.Start()
}

func parseCommandLineArguments() configuration.Configuration {

	contentDirectoryPath := flag.String("content", "data/content", "The path of directory that has to the content to share.")
	logFilePath := flag.String("log", "data/log.txt", "The path of the log file.")
	port := flag.Int("port", 8000, "The port to listen on.")
	useBuiltInServer := flag.Bool("builtin", false, "Use Go's built-in static server.")

	flag.Parse()

	config := configuration.Configuration{*contentDirectoryPath, *logFilePath, *port, *useBuiltInServer}

	return config
}
