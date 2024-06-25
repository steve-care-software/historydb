package jsons

type database struct {
	Head     string   `json:"head"`
	MetaData metaData `json:"meta_data"`
}

type metaData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
