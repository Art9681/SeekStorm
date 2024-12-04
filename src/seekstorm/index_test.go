package seekstorm

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	indexPath := "./test_index"
	index, err := createIndex(indexPath)
	if err != nil {
		t.Fatalf("Failed to create index: %v", err)
	}
	defer os.RemoveAll(indexPath)

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("Index directory was not created: %s", indexPath)
	}

	indexFilePath := filepath.Join(indexPath, "index.json")
	if _, err := os.Stat(indexFilePath); os.IsNotExist(err) {
		t.Errorf("Index file was not created: %s", indexFilePath)
	}
}

func TestOpenIndex(t *testing.T) {
	indexPath := "./test_index"
	_, err := createIndex(indexPath)
	if err != nil {
		t.Fatalf("Failed to create index: %v", err)
	}
	defer os.RemoveAll(indexPath)

	index, err := openIndex(indexPath)
	if err != nil {
		t.Fatalf("Failed to open index: %v", err)
	}

	if index.IndexPath != indexPath {
		t.Errorf("Expected IndexPath to be %s, got %s", indexPath, index.IndexPath)
	}
}

func TestIndexDocument(t *testing.T) {
	indexPath := "./test_index"
	index, err := createIndex(indexPath)
	if err != nil {
		t.Fatalf("Failed to create index: %v", err)
	}
	defer os.RemoveAll(indexPath)

	docID := "doc1"
	document := map[string]interface{}{
		"title": "Test Document",
		"body":  "This is a test document.",
	}

	err = index.indexDocument(docID, document)
	if err != nil {
		t.Fatalf("Failed to index document: %v", err)
	}

	indexFilePath := filepath.Join(indexPath, "index.json")
	file, err := os.Open(indexFilePath)
	if err != nil {
		t.Fatalf("Failed to open index file: %v", err)
	}
	defer file.Close()

	var indexData map[string]interface{}
	if err := json.NewDecoder(file).Decode(&indexData); err != nil {
		t.Fatalf("Failed to read index data: %v", err)
	}

	if _, ok := indexData[docID]; !ok {
		t.Errorf("Document with ID %s was not indexed", docID)
	}
}

func TestSearchIndex(t *testing.T) {
	indexPath := "./test_index"
	index, err := createIndex(indexPath)
	if err != nil {
		t.Fatalf("Failed to create index: %v", err)
	}
	defer os.RemoveAll(indexPath)

	docID := "doc1"
	document := map[string]interface{}{
		"title": "Test Document",
		"body":  "This is a test document.",
	}

	err = index.indexDocument(docID, document)
	if err != nil {
		t.Fatalf("Failed to index document: %v", err)
	}

	results, err := index.searchIndex("Test Document")
	if err != nil {
		t.Fatalf("Failed to search index: %v", err)
	}

	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}

	if results[0]["title"] != "Test Document" {
		t.Errorf("Expected title to be 'Test Document', got '%s'", results[0]["title"])
	}
}
