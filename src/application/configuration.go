package application

import (
	"os"
	"strconv"
)

type Config struct {
	RestServerPort uint16
}

func NewConfig() *Config {
	returnValue := &Config{RestServerPort: 8000}
	returnValue.loadFromEnvironment()
	return returnValue
}

func (c *Config) loadFromEnvironment() {
	if serverPort, exists := os.LookupEnv(""); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			c.RestServerPort = uint16(port)
		}
	}
}
