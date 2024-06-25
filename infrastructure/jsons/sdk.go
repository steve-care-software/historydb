package jsons

import (
	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewDatabaseAdapter creates a new database adapter
func NewDatabaseAdapter() databases.Adapter {
	metaDataBuilder := metadatas.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createDatabaseAdapter(
		metaDataBuilder,
		hashAdapter,
	)
}
