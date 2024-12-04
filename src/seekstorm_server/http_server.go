package seekstorm_server

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
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

type SearchRequestObject struct {
	QueryString    string          `json:"query"`
	Offset         int             `json:"offset"`
	Length         int             `json:"length"`
	ResultType     string          `json:"result_type,omitempty"`
	Realtime       bool            `json:"realtime,omitempty"`
	Highlights     []Highlight     `json:"highlights,omitempty"`
	FieldFilter    []string        `json:"field_filter,omitempty"`
	Fields         []string        `json:"fields,omitempty"`
	DistanceFields []DistanceField `json:"distance_fields,omitempty"`
	QueryFacets    []QueryFacet    `json:"query_facets,omitempty"`
	FacetFilter    []FacetFilter   `json:"facet_filter,omitempty"`
	ResultSort     []ResultSort    `json:"result_sort,omitempty"`
	QueryType      string          `json:"query_type_default,omitempty"`
}

type SearchResultObject struct {
	Time        int64                `json:"time"`
	Query       string               `json:"query"`
	Offset      int                  `json:"offset"`
	Length      int                  `json:"length"`
	Count       int                  `json:"count"`
	CountTotal  int                  `json:"count_total"`
	QueryTerms  []string             `json:"query_terms"`
	Results     []Document           `json:"results"`
	Facets      map[string]Facet     `json:"facets"`
	Suggestions []string             `json:"suggestions"`
}

type CreateIndexRequest struct {
	IndexName  string         `json:"index_name"`
	Schema     []SchemaField  `json:"schema,omitempty"`
	Similarity string         `json:"similarity,omitempty"`
	Tokenizer  string         `json:"tokenizer,omitempty"`
	Synonyms   []Synonym      `json:"synonyms,omitempty"`
}

type DeleteApikeyRequest struct {
	ApikeyBase64 string `json:"apikey_base64"`
}

type GetDocumentRequest struct {
	QueryTerms    []string        `json:"query_terms,omitempty"`
	Highlights    []Highlight     `json:"highlights,omitempty"`
	Fields        []string        `json:"fields,omitempty"`
	DistanceFields []DistanceField `json:"distance_fields,omitempty"`
}

func calculateHash(data []byte) uint64 {
	h := fnv.New64a()
	h.Write(data)
	return h.Sum64()
}

func status(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(errorMessage))
}

func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the HTTP request handler logic here
}

func httpServer(indexPath string, apikeyList *sync.Map, localIP string, localPort int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpRequestHandler(w, r)
	})

	addr := fmt.Sprintf("%s:%d", localIP, localPort)
	fmt.Printf("Listening on: %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("server error: %v\n", err)
		os.Exit(1)
	}
}
