package pointers

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	head        hash.Hash
	metaData    metadatas.MetaData
	path        []string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		head:        nil,
		metaData:    nil,
		path:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head hash.Hash) Builder {
	app.head = head
	return app
}

// WithMetaData adds metaData to the builder
func (app *builder) WithMetaData(metaData metadatas.MetaData) Builder {
	app.metaData = metaData
	return app
}

// WithPath adds path to the builder
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// Now builds a new Pointer instance
func (app *builder) Now() (Pointer, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Pointer instance")
	}

	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Pointer instance")
	}

	if app.metaData == nil {
		return nil, errors.New("the metaData is mandatory in order to build a Pointer instance")
	}

	filePath := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.head.Bytes(),
		app.metaData.Hash().Bytes(),
		[]byte(filePath),
	})

	if err != nil {
		return nil, err
	}

	return createPointer(*pHash, app.head, app.metaData, app.path), nil
}
