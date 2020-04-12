package config

import (
	"log"

	"github.com/tkanos/gonfig"
)

// Configuration - Object to hold all config values
type Configuration struct {
	MongoDBConnectionString string
	MongoDBDatabase         string
}

// GetConfiguration - Retrieves all config values from json fiile
func GetConfiguration() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf("./config/env.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}

	return configuration
}
