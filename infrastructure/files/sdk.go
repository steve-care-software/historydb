package files

import "github.com/steve-care-software/historydb/domain/files"

// NewRepsoitory creates a new reposiotry
func NewRepsoitory(basePath []string) files.Repository {
	return createFileRepository(
		basePath,
	)
}

// NewService creates a new service
func NewService(basePath []string) files.Service {
	return createFileService(
		basePath,
	)
}
