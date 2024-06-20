package chunks

import "github.com/steve-care-software/historydb/domain/hash"

type chunk struct {
	hash        hash.Hash
	path        []string
	fingerPrint hash.Hash
}

func createChunk(
	hash hash.Hash,
	path []string,
	fingerPrint hash.Hash,
) Chunk {
	return &chunk{
		hash:        hash,
		path:        path,
		fingerPrint: fingerPrint,
	}
}

func (obj *chunk) Hash() hash.Hash {
	return obj.hash
}

func (obj *chunk) Path() []string {
	return obj.path
}

func (obj *chunk) FingerPrint() hash.Hash {
	return obj.fingerPrint
}
