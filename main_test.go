package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"testing"
)

func TestLoadData(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write some test data to the file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("user1,product1\nuser2,product2\nuser1,product3\n")
	if err != nil {
		t.Fatal(err)
	}
	writer.Flush()
	file.Close()

	// Create a mock logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Call loadData
	data := loadData(logger, file.Name())

	// Check the results
	if len(data) != 2 {
		t.Errorf("Expected 2 users, got %d", len(data))
	}
	if len(data["user1"]) != 2 {
		t.Errorf("Expected user1 to have 2 products, got %d", len(data["user1"]))
	}
	if len(data["user2"]) != 1 {
		t.Errorf("Expected user2 to have 1 product, got %d", len(data["user2"]))
	}
	if !data["user1"]["product1"] {
		t.Errorf("Expected user1 to have product1")
	}
	if !data["user1"]["product3"] {
		t.Errorf("Expected user1 to have product3")
	}
	if !data["user2"]["product2"] {
		t.Errorf("Expected user2 to have product2")
	}
}

func TestFindNewPagesVisited(t *testing.T) {
	// Create a temporary file for the first day
	firstDayFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(firstDayFile.Name())

	// Write some test data to the first day file
	writer := bufio.NewWriter(firstDayFile)
	_, err = writer.WriteString("user1,product1\nuser2,product2\nuser1,product3\n")
	if err != nil {
		t.Fatal(err)
	}
	writer.Flush()
	firstDayFile.Close()

	// Create a temporary file for the second day
	secondDayFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(secondDayFile.Name())

	// Write some test data to the second day file
	writer = bufio.NewWriter(secondDayFile)
	_, err = writer.WriteString("user1,product4\nuser2,product2\nuser1,product5\n")
	if err != nil {
		t.Fatal(err)
	}
	writer.Flush()
	secondDayFile.Close()

	// Create a mock logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Load the first day data
	firstDayData := loadData(logger, firstDayFile.Name())

	// Capture the standard output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call findNewPagesVisited
	findNewPagesVisited(logger, secondDayFile.Name(), firstDayData)

	// Restore the standard output
	w.Close()
	os.Stdout = oldStdout

	// Read the captured standard output
	out, _ := io.ReadAll(r)

	// Check the results
	expected := "user1\nuser1\n"
	if string(out) != expected {
		t.Errorf("Expected output %q, got %q", expected, string(out))
	}
}
