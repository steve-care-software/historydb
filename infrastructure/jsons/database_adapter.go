package jsons

import (
	"encoding/json"

	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/hash"
)

type databaseAdapter struct {
	metaDataBuilder metadatas.Builder
	hashAdapter     hash.Adapter
}

func createDatabaseAdapter(
	metaDataBuilder metadatas.Builder,
	hashAdapter hash.Adapter,
) databases.Adapter {
	out := databaseAdapter{
		metaDataBuilder: metaDataBuilder,
		hashAdapter:     hashAdapter,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *databaseAdapter) ToBytes(ins databases.Database) ([]byte, error) {
	metaDataIns := ins.MetaData()
	return json.Marshal(database{
		Head: ins.Hash().String(),
		MetaData: metaData{
			Name:        metaDataIns.Name(),
			Description: metaDataIns.Description(),
		},
	})
}

// ToComponents bytes to components
func (app *databaseAdapter) ToComponents(bytes []byte, path []string) (metadatas.MetaData, hash.Hash, error) {
	ptr := new(database)
	err := json.Unmarshal(bytes, ptr)
	if err != nil {
		return nil, nil, err
	}

	pHeadHash, err := app.hashAdapter.FromString(ptr.Head)
	if err != nil {
		return nil, nil, err
	}

	metaDataStr := ptr.MetaData
	metaData, err := app.metaDataBuilder.Create().
		WithPath(path).
		WithName(metaDataStr.Name).
		WithDescription(metaDataStr.Description).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return metaData, *pHeadHash, nil
}
