package executions

import (
	"github.com/steve-care-software/historydb/domain/databases/commits/executions/chunks"
	"github.com/steve-care-software/historydb/domain/hash"
)

// Executions represents executions
type Executions interface {
	Hash() hash.Hash
	List() []Execution
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	IsBytes() bool
	Bytes() []byte
	IsChunk() bool
	Chunk() chunks.Chunk
}
