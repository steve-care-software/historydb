package chunks

import "github.com/steve-care-software/historydb/domain/hash"

// Chunk represents a chunk
type Chunk interface {
	Hash() hash.Hash
	Path() []string
	FingerPrint() hash.Hash
}
