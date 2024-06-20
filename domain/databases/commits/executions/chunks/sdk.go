package chunks

import "github.com/steve-care-software/historydb/domain/hash"

// NewChunkBuilder creates a new chunk builder
func NewChunkBuilder() ChunkBuilder {
	hashAdapter := hash.NewAdapter()
	return createChunkBuilder(
		hashAdapter,
	)
}

// ChunkBuilder represents the chunk builder
type ChunkBuilder interface {
	Create() ChunkBuilder
	WithPath(path []string) ChunkBuilder
	WithFingerPrint(fingerPrint hash.Hash) ChunkBuilder
	Now() (Chunk, error)
}

// Chunk represents a chunk
type Chunk interface {
	Hash() hash.Hash
	Path() []string
	FingerPrint() hash.Hash
}
