package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	_ "modernc.org/sqlite"
)

// Config struct to hold the history path
type Config struct {
	HistoryPath string `yaml:"history_path"`
}

// Function to read configuration from file
func readConfig(filePath string) (Config, error) {
	var config Config
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func main() {
	// Path to the configuration file
	configPath := "config.yml"

	// Read configuration
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Validate if the file exists and is accessible
	if _, err := os.Stat(config.HistoryPath); os.IsNotExist(err) {
		log.Fatalf("History file does not exist: %v", err)
	} else if err != nil {
		log.Fatalf("Error accessing history file: %v", err)
	}

	// Open the SQLite database
	db, err := sql.Open("sqlite", config.HistoryPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Query the database for browser history
	rows, err := db.Query("SELECT url, title, visit_count, last_visit_time FROM urls ORDER BY last_visit_time DESC")
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}
	defer rows.Close()

	// Print the history
	for rows.Next() {
		var url, title string
		var visitCount int
		var lastVisitTime int64
		err := rows.Scan(&url, &title, &visitCount, &lastVisitTime)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		fmt.Printf("URL: %s\nTitle: %s\nVisit Count: %d\nLast Visit Time: %d\n\n", url, title, visitCount, lastVisitTime)
	}

	// Check for errors from iterating over rows.
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}
}
