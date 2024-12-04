package seekstorm_server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	data := []byte("test data")
	expectedHash := uint64(0x3a0b0c8e)
	hash := calculateHash(data)
	if hash != expectedHash {
		t.Errorf("Expected hash %x, but got %x", expectedHash, hash)
	}
}

func TestStatus(t *testing.T) {
	rr := httptest.NewRecorder()
	status(rr, http.StatusNotFound, "Not Found")
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, rr.Code)
	}
	if rr.Body.String() != "Not Found" {
		t.Errorf("Expected body 'Not Found', but got '%s'", rr.Body.String())
	}
}

func TestHttpRequestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpRequestHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"message": "Hello, World!"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHttpServer(t *testing.T) {
	go func() {
		httpServer("./test_index", &sync.Map{}, "127.0.0.1", 8080)
	}()
	time.Sleep(1 * time.Second)
	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}
