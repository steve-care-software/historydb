package metadatas

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

// MetaData represents a database metadata
type MetaData interface {
	Hash() hash.Hash
	Path() []string
	Name() string
	Description() string
}
