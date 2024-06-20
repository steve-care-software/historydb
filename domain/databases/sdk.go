package databases

import (
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new database builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithHead(head commits.Commit) Builder
	WithMetaData(metaData metadatas.MetaData) Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Hash() hash.Hash
	Head() commits.Commit
	MetaData() metadatas.MetaData
}
