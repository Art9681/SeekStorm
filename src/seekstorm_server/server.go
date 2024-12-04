package seekstorm_server

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
)

type ApikeyObject struct {
	ID         uint64
	ApikeyHash uint64
	Quota      ApikeyQuotaObject
	IndexList  map[uint64]*Index
}

type ApikeyQuotaObject struct {
	IndicesMax     int
	IndicesSizeMax int
	DocumentsMax   int
	OperationsMax  int
	RateLimit      int
}

type Index struct {
	// Add appropriate fields for the Index struct
}

func initialize(params map[string]string) {
	ingestPath := params["ingest_path"]
	if ingestPath == "" {
		ingestPath = "./"
	}
	fmt.Printf("Ingest path: %s\n", ingestPath)

	indexPath := params["index_path"]
	if indexPath == "" {
		indexPath = "./seekstorm_index"
	}

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		if err := os.MkdirAll(indexPath, os.ModePerm); err != nil {
			fmt.Printf("index_path could not be created: %s\n", indexPath)
		} else {
			fmt.Printf("index_path did not exist, new directory created: %s\n", indexPath)
		}
	}

	apikeyList := sync.Map{}
	openAllApikeys(indexPath, &apikeyList)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		httpServer(indexPath, &apikeyList, "0.0.0.0", 80)
	}()

	go func() {
		commandline()
	}()

	<-sigChan
	fmt.Println("Committing all indices ...")
	commitAllIndices(&apikeyList)
	fmt.Println("Server stopped by signal")
}

func commandline() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch strings.ToLower(input) {
		case "quit":
			fmt.Println("Committing all indices ...")
			commitAllIndices(nil)
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

func ctrlChannel() chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	return sigChan
}

func commitAllIndices(apikeyList *sync.Map) {
	// Implement the logic to commit all indices
}

func openAllApikeys(indexPath string, apikeyList *sync.Map) {
	// Implement the logic to open all API keys
}

func httpServer(indexPath string, apikeyList *sync.Map, localIP string, localPort int) {
	// Implement the HTTP server logic
}
