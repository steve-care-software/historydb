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

// Adapter represents a database adapter
type Adapter interface {
	ToBytes(ins Database) ([]byte, error)
	ToInstance(bytes []byte) (Database, error)
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

// Repository represents a database repository
type Repository interface {
	Retrieve(path []string) (Database, error)
}

// Service represents a database service
type Service interface {
	Save(database Database) error
	SaveAll(list []Database) error
}
