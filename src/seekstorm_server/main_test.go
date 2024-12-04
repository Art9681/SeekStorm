package main

import (
	"testing"
	"os"
	"os/signal"
	"syscall"
)

func TestMainFunction(t *testing.T) {
	// Set up environment variable for testing
	os.Setenv("MASTER_KEY_SECRET", "test_master_key")

	// Run the main function in a separate goroutine
	go main()

	// Simulate a shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sigChan <- syscall.SIGINT

	// Wait for the main function to complete
	<-sigChan

	// Check if the server stopped gracefully
	// You can add more assertions here if needed
}

func TestInitializeFunction(t *testing.T) {
	// Set up test parameters
	indexPath := "./test_index"
	localIP := "127.0.0.1"
	localPort := 8080
	masterKey := "test_master_key"

	// Run the initialize function
	initialize(indexPath, localIP, localPort, masterKey)

	// Check if the index path was created
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("index path was not created: %s", indexPath)
	}

	// Clean up test index path
	os.RemoveAll(indexPath)
}
