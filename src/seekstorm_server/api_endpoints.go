package seekstorm_server

import (
	"encoding/json"
)

type SearchRequestObject struct {
	QueryString    string          `json:"query"`
	Offset         int             `json:"offset"`
	Length         int             `json:"length"`
	ResultType     ResultType      `json:"result_type,omitempty"`
	Realtime       bool            `json:"realtime,omitempty"`
	Highlights     []Highlight     `json:"highlights,omitempty"`
	FieldFilter    []string        `json:"field_filter,omitempty"`
	Fields         []string        `json:"fields,omitempty"`
	DistanceFields []DistanceField `json:"distance_fields,omitempty"`
	QueryFacets    []QueryFacet    `json:"query_facets,omitempty"`
	FacetFilter    []FacetFilter   `json:"facet_filter,omitempty"`
	ResultSort     []ResultSort    `json:"result_sort,omitempty"`
	QueryType      QueryType       `json:"query_type_default,omitempty"`
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
	IndexName string         `json:"index_name"`
	Schema    []SchemaField  `json:"schema,omitempty"`
	Similarity SimilarityType `json:"similarity,omitempty"`
	Tokenizer TokenizerType  `json:"tokenizer,omitempty"`
	Synonyms  []Synonym      `json:"synonyms,omitempty"`
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
