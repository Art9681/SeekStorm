package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var (
	indexPath   string
	localIP     string
	localPort   int
	masterKey   string
)

func init() {
	flag.StringVar(&indexPath, "index_path", "./seekstorm_index", "Path to the index directory")
	flag.StringVar(&localIP, "local_ip", "0.0.0.0", "Local IP address to bind the server")
	flag.IntVar(&localPort, "local_port", 80, "Local port to bind the server")
	flag.StringVar(&masterKey, "master_key", "", "Master key for API key management")
}

func main() {
	flag.Parse()

	if masterKey == "" {
		fmt.Println("MASTER_KEY_SECRET environment variable is not set")
		os.Exit(1)
	}

	fmt.Printf("SeekStorm server v%s\n", Version)
	fmt.Println("Press CTRL-C or enter 'quit' to shutdown server, enter 'help' for console commands.")

	initialize(indexPath, localIP, localPort, masterKey)
}

func initialize(indexPath, localIP string, localPort int, masterKey string) {
	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the HTTP server
	go func() {
		httpServer(indexPath, localIP, localPort)
	}()

	// Command line input handling
	go func() {
		commandline()
	}()

	// Wait for shutdown signal
	<-sigChan
	fmt.Println("Committing all indices ...")
	commitAllIndices()
	fmt.Println("Server stopped by signal")
}

func commandline() {
	for {
		var input string
		fmt.Scanln(&input)
		switch strings.ToLower(input) {
		case "quit":
			fmt.Println("Committing all indices ...")
			commitAllIndices()
			fmt.Println("Server stopped by quit")
			os.Exit(0)
		case "help":
			fmt.Println("Server console commands:")
			fmt.Println("  quit  - Stop the server")
			fmt.Println("  help  - Show this help message")
		default:
			fmt.Printf("Unknown command: %s\n", input)
		}
	}
}

func commitAllIndices() {
	// Implement the logic to commit all indices
}
