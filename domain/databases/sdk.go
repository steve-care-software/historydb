package databases

import (
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/hash"
)

// Database represents a database
type Database interface {
	Hash() hash.Hash
	Head() commits.Commit
	MetaData() metadatas.MetaData
}
