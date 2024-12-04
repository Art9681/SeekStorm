package seekstorm_server

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestInitialize(t *testing.T) {
	params := map[string]string{
		"ingest_path": "./",
		"index_path":  "./seekstorm_index",
	}

	initialize(params)
}

func TestCommandline(t *testing.T) {
	go func() {
		commandline()
	}()

	// Simulate user input
	input := "quit\n"
	os.Stdin.Write([]byte(input))
}

func TestCtrlChannel(t *testing.T) {
	sigChan := ctrlChannel()

	// Send a signal to the channel
	signal.Notify(sigChan, syscall.SIGINT)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)

	// Wait for the signal
	<-sigChan
}
