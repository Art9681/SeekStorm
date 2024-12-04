package seekstorm

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Index struct {
	IndexPath   string
	IndexData   map[string]interface{}
	IndexMutex  sync.RWMutex
}

func createIndex(indexPath string) (*Index, error) {
	index := &Index{
		IndexPath: indexPath,
		IndexData: make(map[string]interface{}),
	}

	if err := os.MkdirAll(indexPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create index directory: %w", err)
	}

	indexFilePath := filepath.Join(indexPath, "index.json")
	file, err := os.Create(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create index file: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(index.IndexData); err != nil {
		return nil, fmt.Errorf("failed to write index data: %w", err)
	}

	return index, nil
}

func openIndex(indexPath string) (*Index, error) {
	index := &Index{
		IndexPath: indexPath,
		IndexData: make(map[string]interface{}),
	}

	indexFilePath := filepath.Join(indexPath, "index.json")
	file, err := os.Open(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open index file: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&index.IndexData); err != nil {
		return nil, fmt.Errorf("failed to read index data: %w", err)
	}

	return index, nil
}

func (index *Index) indexDocument(docID string, document map[string]interface{}) error {
	index.IndexMutex.Lock()
	defer index.IndexMutex.Unlock()

	index.IndexData[docID] = document

	indexFilePath := filepath.Join(index.IndexPath, "index.json")
	file, err := os.Create(indexFilePath)
	if err != nil {
		return fmt.Errorf("failed to create index file: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(index.IndexData); err != nil {
		return fmt.Errorf("failed to write index data: %w", err)
	}

	return nil
}

func (index *Index) searchIndex(query string) ([]map[string]interface{}, error) {
	index.IndexMutex.RLock()
	defer index.IndexMutex.RUnlock()

	results := []map[string]interface{}{}
	for _, doc := range index.IndexData {
		document, ok := doc.(map[string]interface{})
		if !ok {
			continue
		}

		for _, value := range document {
			if strValue, ok := value.(string); ok && strValue == query {
				results = append(results, document)
				break
			}
		}
	}

	if len(results) == 0 {
		return nil, errors.New("no documents found")
	}

	return results, nil
}
