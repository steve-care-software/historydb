package chunks

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/historydb/domain/hash"
)

type chunkBuilder struct {
	hashAdapter hash.Adapter
	path        []string
	fingerPrint hash.Hash
}

func createChunkBuilder(
	hashAdapter hash.Adapter,
) ChunkBuilder {
	return &chunkBuilder{
		hashAdapter: hashAdapter,
		path:        nil,
		fingerPrint: nil,
	}
}

// Create initializes the builder
func (app *chunkBuilder) Create() ChunkBuilder {
	return createChunkBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the buiolder
func (app *chunkBuilder) WithPath(path []string) ChunkBuilder {
	app.path = path
	return app
}

// WithFingerPrint adds a fingerprint to the builder
func (app *chunkBuilder) WithFingerPrint(fingerPrint hash.Hash) ChunkBuilder {
	app.fingerPrint = fingerPrint
	return app
}

// Now builds a new Chunk instance
func (app *chunkBuilder) Now() (Chunk, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Chunk instance")
	}

	if app.fingerPrint == nil {
		return nil, errors.New("the fingerPrint is mandatory in order to build a Chunk instance")
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(path),
		app.fingerPrint.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createChunk(
		*pHash,
		app.path,
		app.fingerPrint,
	), nil
}
