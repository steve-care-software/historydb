package commits

import (
	"github.com/steve-care-software/historydb/domain/databases/commits/executions"
	"github.com/steve-care-software/historydb/domain/hash"
)

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Executions() executions.Executions
	HasParent() bool
	Parent() hash.Hash
}
