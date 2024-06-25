package databases

import (
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/files"
)

type service struct {
	fileService     files.Service
	commitService   commits.Service
	databaseAdapter Adapter
}

func createService(
	fileService files.Service,
	commitService commits.Service,
	databaseAdapter Adapter,
) Service {
	out := service{
		fileService:     fileService,
		commitService:   commitService,
		databaseAdapter: databaseAdapter,
	}

	return &out
}

// Save saves a database
func (app *service) Save(ins Database) error {
	err := app.commitService.Save(ins.Head())
	if err != nil {
		return err
	}

	bytes, err := app.databaseAdapter.ToBytes(ins)
	if err != nil {
		return err
	}

	path := ins.MetaData().Path()
	return app.fileService.Save(path, bytes)
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
