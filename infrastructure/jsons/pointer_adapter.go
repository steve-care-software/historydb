package jsons

import (
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/databases/pointers"
	"github.com/steve-care-software/historydb/domain/hash"
)

type pointerAdapter struct {
	metaDataBuilder metadatas.Builder
	hashAdapter     hash.Adapter
}

func createPointerAdapter(
	metaDataBuilder metadatas.Builder,
	hashAdapter hash.Adapter,
) pointers.Adapter {
	out := pointerAdapter{
		metaDataBuilder: metaDataBuilder,
		hashAdapter:     hashAdapter,
	}

	return &out
}

// ToBytes converts pointer to bytes
func (app *pointerAdapter) ToBytes(ins pointers.Pointer) ([]byte, error) {
	return nil, nil
}

// ToInstance converts bytes to pointer
func (app *pointerAdapter) ToInstance(bytes []byte) (pointers.Pointer, error) {
	return nil, nil
}
