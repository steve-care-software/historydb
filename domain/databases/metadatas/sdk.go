package metadatas

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new metadata builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a metadata builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	Now() (MetaData, error)
}

// MetaData represents a database metadata
type MetaData interface {
	Hash() hash.Hash
	Path() []string
	Name() string
	Description() string
}
