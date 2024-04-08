package main

import "fmt"

// Config represents the application configuration
type Config struct {
	LogLevel string
	Port     int
}

// UpdateConfig modifies the provided configuration
func UpdateConfig(c *Config, logLevel string, port int) {
	c.LogLevel = logLevel
	c.Port = port
}

func main() {
	// Initial configuration
	appConfig := &Config{
		LogLevel: "info",
		Port:     8080,
	}

	fmt.Println("Initial Config:", appConfig)

	// Update configuration
	UpdateConfig(appConfig, "debug", 9000)
	fmt.Println("Updated Config:", appConfig)
}