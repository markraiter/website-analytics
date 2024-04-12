package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func main() {
	// Check if the correct number of arguments are passed
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./main day1.csv day2.csv")
		os.Exit(1)
	}

	// Initialize logger
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Load data from the first day into a map where the key is the user ID and the value is a set of visited product IDs
	firstDayData := loadData(log, os.Args[1])

	// Load data from the second day and find users who visited new pages
	findNewPagesVisited(log, os.Args[2], firstDayData)
}

// loadData loads data from a CSV file and returns a map where the key is the user ID and the value is a set of visited product IDs
func loadData(log *slog.Logger, filename string) map[string]map[string]bool {
	file, err := os.Open(filename)
	if err != nil {
		log.Error("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	data := make(map[string]map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		userID := fields[0]
		productID := fields[1]

		if _, ok := data[userID]; !ok {
			data[userID] = make(map[string]bool)
		}
		data[userID][productID] = true
	}

	if err := scanner.Err(); err != nil {
		log.Error("Error reading file:", err)
		os.Exit(1)
	}

	return data
}

// findNewPagesVisited finds users who visited new pages on the second day compared to the first day
func findNewPagesVisited(log *slog.Logger, filename string, firstDayData map[string]map[string]bool) {
	file, err := os.Open(filename)
	if err != nil {
		log.Error("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		userID := fields[0]
		productID := fields[1]

		if visitedPages, ok := firstDayData[userID]; ok {
			if _, visited := visitedPages[productID]; !visited {
				fmt.Println(userID)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("Error reading file:", err)
		os.Exit(1)
	}
}
