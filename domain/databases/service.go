package databases

import (
	"github.com/steve-care-software/historydb/domain/files"
)

type service struct {
	fileService     files.Service
	databaseAdapter Adapter
}

func createService(
	fileService files.Service,
	databaseAdapter Adapter,
) Service {
	out := service{
		fileService:     fileService,
		databaseAdapter: databaseAdapter,
	}

	return &out
}

// Save saves a database
func (app *service) Save(ins Database) error {
	bytes, err := app.databaseAdapter.ToBytes(ins)
	if err != nil {
		return err
	}

	path := ins.MetaData().Path()
	return app.fileService.Transact(path, bytes)
}

// SaveAll saves all databases
func (app *service) SaveAll(list []Database) error {
	for _, oneIns := range list {
		err := app.Save(oneIns)
		if err != nil {
			return err
		}
	}

	return nil
}
