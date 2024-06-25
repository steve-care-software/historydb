package files

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/historydb/domain/files"
)

type fileService struct {
	basePath []string
	locks    map[string]*fslock.Lock
}

func createFileService(
	basePath []string,
) files.Service {
	out := fileService{
		locks: map[string]*fslock.Lock{},
	}

	return &out
}

// Init initializes a file
func (app *fileService) Init(path []string) error {
	filePath := createFilePath(app.basePath, path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.Mkdir(filePath, os.ModeDir)
		if err != nil {
			return err
		}
	}

	return nil
}

// Lock locks a file
func (app *fileService) Lock(path []string) error {
	filePath := createFilePath(app.basePath, path)
	lock := fslock.New(filePath)
	err := lock.TryLock()
	if err != nil {
		str := fmt.Sprintf("failed to acquire lock: %s", err.Error())
		return errors.New(str)
	}

	app.locks[filePath] = lock
	return nil
}

// Unlock unlocks a file
func (app *fileService) Unlock(path []string) error {
	filePath := createFilePath(app.basePath, path)
	if pLock, ok := app.locks[filePath]; ok {
		pLock.Unlock()
		return nil
	}

	str := fmt.Sprintf("there is no lock on the provided file path: %s", filePath)
	return errors.New(str)
}

// Save saves data in a file
func (app *fileService) Save(path []string, bytes []byte) error {
	filePath := createFilePath(app.basePath, path)
	return ioutil.WriteFile(filePath, bytes, fs.ModePerm)
}
