package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Config represents the application configuration.
func InitConfig() {
	configFile, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("error opening config file: ", err)
	}
	defer configFile.Close()
	var config map[string]string
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatal("error decoding config file: ", err)
	}
	// Initialize database connection
	dbUser = config["db_user"]
	dbPassword = config["db_password"]
	dbName = config["db_name"]
	dbHost = config["db_host"]
	dbPort = config["db_port"]
}