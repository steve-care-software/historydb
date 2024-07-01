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
		"dbfile.data",
	}

	expectedBytes := []byte("this is some data")

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	databaseBasePath := append(basePath, "databases", "my_database")
	commitBasePath := []string{"commits"}
	chunksInnerPath := []string{"chunks"}

	chunkFileRepository, err := files.NewRepositoryBuilder(chunksInnerPath).Create().WithBasePath(databaseBasePath).Now()
	if err != nil {
		panic(err)
	}

	chunkFileService, err := files.NewServiceBuilder(chunksInnerPath).Create().WithBasePath(databaseBasePath).Now()
	if err != nil {
		panic(err)
	}

	databaseRepository, databaseService, commitRepository, _ := bundles.NewDatabaseRepositoryServiceWithJsonAdapter(databaseBasePath, commitBasePath)
	application := NewApplication(
		databaseRepository,
		databaseService,
		commitRepository,
		chunkFileRepository,
		chunkFileService,
		uint(1024),
	)

	// init, begin, commit once, push
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

	retDatabase, err := application.Retrieve(path)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retDatabase.MetaData().Name() != name {
		t.Errorf("the metaData name was expected to be '%s', '%s' returned", name, retDatabase.MetaData().Name())
		return
	}

	if retDatabase.MetaData().Description() != description {
		t.Errorf("the metaData description was expected to be '%s', '%s' returned", description, retDatabase.MetaData().Description())
		return
	}
}
