package databases

import (
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/files"
)

type repository struct {
	fileRepository   files.Repository
	commitRepository commits.Repository
	databaseAdapter  Adapter
	databaseBuilder  Builder
}

func createRepository(
	fileRepository files.Repository,
	commitRepository commits.Repository,
	databaseAdapter Adapter,
	databaseBuilder Builder,
) Repository {
	out := repository{
		fileRepository:   fileRepository,
		commitRepository: commitRepository,
		databaseAdapter:  databaseAdapter,
		databaseBuilder:  databaseBuilder,
	}

	return &out
}

// Retrieve retrieves a database by path
func (app *repository) Retrieve(path []string) (Database, error) {
	bytes, err := app.fileRepository.Retrieve(path)
	if err != nil {
		return nil, err
	}

	metaData, headHash, err := app.databaseAdapter.ToComponents(bytes, path)
	if err != nil {
		return nil, err
	}

	head, err := app.commitRepository.Retrieve(headHash)
	if err != nil {
		return nil, err
	}

	return app.databaseBuilder.Create().
		WithHead(head).
		WithMetaData(metaData).
		Now()
}
