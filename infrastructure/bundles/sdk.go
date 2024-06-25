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
) (commits.Repository, commits.Service) {
	commitFileRepository := files.NewRepsoitory(basePath)
	commitAdapter := jsons.NewCommitAdapter()
	commitRepository := commits.NewRepository(
		commitAdapter,
		commitFileRepository,
	)

	commitFileService := files.NewService(
		commitFileRepository,
		basePath,
	)

	commitService := commits.NewService(
		commitAdapter,
		commitFileService,
	)

	return commitRepository, commitService
}

// NewDatabaseRepositoryServiceWithJsonAdapter creates a new database repository and service with json adapter
func NewDatabaseRepositoryServiceWithJsonAdapter(
	commitBasePath []string,
	databaseBasePath []string,
) (databases.Repository, databases.Service, commits.Repository, commits.Service) {
	commitRepository, commitService := NewCommitRepositoryServiceWithJsonAdapter(commitBasePath)
	databaseFileRepository := files.NewRepsoitory(databaseBasePath)
	databaseAdapter := jsons.NewDatabaseAdapter()
	databaseRepository := databases.NewRepository(
		databaseFileRepository,
		commitRepository,
		databaseAdapter,
	)

	databaseFileService := files.NewService(
		databaseFileRepository,
		databaseBasePath,
	)

	databaseService := databases.NewService(
		databaseFileService,
		commitService,
		databaseAdapter,
	)

	return databaseRepository, databaseService, commitRepository, commitService
}
