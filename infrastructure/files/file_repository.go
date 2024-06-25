package files

import (
	"os"

	"github.com/steve-care-software/historydb/domain/files"
)

type fileRepository struct {
	basePath []string
}

func createFileRepository(
	basePath []string,
) files.Repository {
	out := fileRepository{
		basePath: basePath,
	}

	return &out
}

// Exists returns true if the file exists, false otherwise
func (app *fileRepository) Exists(path []string) bool {
	filePath := createFilePath(app.basePath, path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}
