package seekstorm_server

import (
	"encoding/json"
	"testing"
)

func TestSearchRequestObject(t *testing.T) {
	jsonData := `{"query": "test query", "offset": 0, "length": 10}`
	var obj SearchRequestObject
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if obj.QueryString != "test query" {
		t.Errorf("Expected QueryString to be 'test query', got '%s'", obj.QueryString)
	}
	if obj.Offset != 0 {
		t.Errorf("Expected Offset to be 0, got %d", obj.Offset)
	}
	if obj.Length != 10 {
		t.Errorf("Expected Length to be 10, got %d", obj.Length)
	}
}

func TestSearchResultObject(t *testing.T) {
	jsonData := `{"time": 123456789, "query": "test query", "offset": 0, "length": 10, "count": 1, "count_total": 1, "query_terms": ["test"], "results": [], "facets": {}, "suggestions": []}`
	var obj SearchResultObject
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if obj.Time != 123456789 {
		t.Errorf("Expected Time to be 123456789, got %d", obj.Time)
	}
	if obj.Query != "test query" {
		t.Errorf("Expected Query to be 'test query', got '%s'", obj.Query)
	}
	if obj.Offset != 0 {
		t.Errorf("Expected Offset to be 0, got %d", obj.Offset)
	}
	if obj.Length != 10 {
		t.Errorf("Expected Length to be 10, got %d", obj.Length)
	}
	if obj.Count != 1 {
		t.Errorf("Expected Count to be 1, got %d", obj.Count)
	}
	if obj.CountTotal != 1 {
		t.Errorf("Expected CountTotal to be 1, got %d", obj.CountTotal)
	}
	if len(obj.QueryTerms) != 1 || obj.QueryTerms[0] != "test" {
		t.Errorf("Expected QueryTerms to be ['test'], got %v", obj.QueryTerms)
	}
	if len(obj.Results) != 0 {
		t.Errorf("Expected Results to be empty, got %v", obj.Results)
	}
	if len(obj.Facets) != 0 {
		t.Errorf("Expected Facets to be empty, got %v", obj.Facets)
	}
	if len(obj.Suggestions) != 0 {
		t.Errorf("Expected Suggestions to be empty, got %v", obj.Suggestions)
	}
}

func TestCreateIndexRequest(t *testing.T) {
	jsonData := `{"index_name": "test_index"}`
	var obj CreateIndexRequest
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if obj.IndexName != "test_index" {
		t.Errorf("Expected IndexName to be 'test_index', got '%s'", obj.IndexName)
	}
}

func TestDeleteApikeyRequest(t *testing.T) {
	jsonData := `{"apikey_base64": "test_apikey"}`
	var obj DeleteApikeyRequest
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if obj.ApikeyBase64 != "test_apikey" {
		t.Errorf("Expected ApikeyBase64 to be 'test_apikey', got '%s'", obj.ApikeyBase64)
	}
}

func TestGetDocumentRequest(t *testing.T) {
	jsonData := `{"query_terms": ["test"], "highlights": [], "fields": ["field1"], "distance_fields": []}`
	var obj GetDocumentRequest
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if len(obj.QueryTerms) != 1 || obj.QueryTerms[0] != "test" {
		t.Errorf("Expected QueryTerms to be ['test'], got %v", obj.QueryTerms)
	}
	if len(obj.Highlights) != 0 {
		t.Errorf("Expected Highlights to be empty, got %v", obj.Highlights)
	}
	if len(obj.Fields) != 1 || obj.Fields[0] != "field1" {
		t.Errorf("Expected Fields to be ['field1'], got %v", obj.Fields)
	}
	if len(obj.DistanceFields) != 0 {
		t.Errorf("Expected DistanceFields to be empty, got %v", obj.DistanceFields)
	}
}
