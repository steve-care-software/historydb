package jsons

// Database represents a database
type Database struct {
	Head     string   `json:"head"`
	MetaData MetaData `json:"meta_data"`
}

// MetaData represents a metadata
type MetaData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
