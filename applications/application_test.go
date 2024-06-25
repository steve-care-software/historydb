package applications

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-care-software/historydb/infrastructure/bundles"
	"github.com/steve-care-software/historydb/infrastructure/files"
)

func TestApplication_beginWithInit_Success(t *testing.T) {
	basePath := []string{
		"test_files",
	}

	path := []string{
		"some_dir",
		"myFile.data",
	}

	expectedBytes := []byte("this is some data")

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	commitBasePath := append(basePath, "commits")
	chunksBasePath := append(basePath, "chunks")
	databaseFileRepository := files.NewRepsoitory(basePath)
	databaseFileService := files.NewService(
		databaseFileRepository,
		basePath,
	)

	databaseRepository, databaseService, commitRepository, _ := bundles.NewDatabaseRepositoryServiceWithJsonAdapter(commitBasePath, basePath)
	application := NewApplication(
		databaseRepository,
		databaseService,
		commitRepository,
		databaseFileRepository,
		databaseFileService,
		chunksBasePath,
		uint(1024),
	)

	name := "My Name"
	description := "This is a description"
	pContext, err := application.BeginWithInit(path, name, description)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Execute(*pContext, expectedBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Push(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}
}
