package commits

import (
	"github.com/steve-care-software/historydb/domain/databases/commits/executions"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new commit builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents a commit adapter
type Adapter interface {
	ToBytes(ins Commit) ([]byte, error)
	ToInstance(bytes []byte) (Commit, error)
}

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithExecutions(executions executions.Executions) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Executions() executions.Executions
	HasParent() bool
	Parent() hash.Hash
}

// Repository represents a commit repository
type Repository interface {
	Retrieve(hash hash.Hash) (Commit, error)
}

// Service represents the commit service
type Service interface {
	Save(ins Commit) error
}
