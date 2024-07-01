package bundles

import (
	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/infrastructure/files"
	"github.com/steve-care-software/historydb/infrastructure/jsons"
)

// NewCommitRepositoryServiceWithJsonAdapter creates a new commit repository and service with json adapter
func NewCommitRepositoryServiceWithJsonAdapter(
	basePath []string,
	innerPath []string,
) (commits.Repository, commits.Service) {
	commitFileRepository, err := files.NewRepositoryBuilder(innerPath).Create().
		WithBasePath(basePath).
		Now()

	if err != nil {
		panic(err)
	}

	commitAdapter := jsons.NewCommitAdapter()
	commitRepository := commits.NewRepository(
		commitAdapter,
		commitFileRepository,
	)

	commitFileService, err := files.NewServiceBuilder(innerPath).Create().
		WithBasePath(basePath).
		Now()

	if err != nil {
		panic(err)
	}

	commitService := commits.NewService(
		commitAdapter,
		commitFileService,
	)

	return commitRepository, commitService
}

// NewDatabaseRepositoryServiceWithJsonAdapter creates a new database repository and service with json adapter
func NewDatabaseRepositoryServiceWithJsonAdapter(
	basePath []string,
	commitInnerPath []string,
) (databases.Repository, databases.Service, commits.Repository, commits.Service) {
	commitRepository, commitService := NewCommitRepositoryServiceWithJsonAdapter(basePath, commitInnerPath)
	databaseFileRepository, err := files.NewRepositoryBuilder([]string{}).Create().
		WithBasePath(basePath).
		Now()

	if err != nil {
		panic(err)
	}

	databaseAdapter := jsons.NewDatabaseAdapter()
	databaseRepository := databases.NewRepository(
		databaseFileRepository,
		commitRepository,
		databaseAdapter,
	)

	databaseFileService, err := files.NewServiceBuilder([]string{}).Create().
		WithBasePath(basePath).
		Now()

	if err != nil {
		panic(err)
	}

	databaseService := databases.NewService(
		databaseRepository,
		databaseFileService,
		commitService,
		databaseAdapter,
	)

	return databaseRepository, databaseService, commitRepository, commitService
}
