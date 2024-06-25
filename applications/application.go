package applications

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/databases/commits/executions"
	"github.com/steve-care-software/historydb/domain/databases/commits/executions/chunks"
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/files"
	"github.com/steve-care-software/historydb/domain/hash"
)

type application struct {
	hashAdapter           hash.Adapter
	repository            databases.Repository
	service               databases.Service
	commitRepository      commits.Repository
	fileRepository        files.Repository
	fileService           files.Service
	databaseBuilder       databases.Builder
	commitBuilder         commits.Builder
	executionsBuilder     executions.Builder
	executionBuilder      executions.ExecutionBuilder
	metaDataBuilder       metadatas.Builder
	chunkBuilder          chunks.Builder
	chunkBasePath         []string
	minSizeToChunkInBytes uint
	commits               map[uint]commit
	contexts              map[uint]contexts
}

func createApplication(
	hashAdapter hash.Adapter,
	repository databases.Repository,
	service databases.Service,
	commitRepository commits.Repository,
	fileRepository files.Repository,
	fileService files.Service,
	databaseBuilder databases.Builder,
	commitBuilder commits.Builder,
	executionsBuilder executions.Builder,
	executionBuilder executions.ExecutionBuilder,
	metaDataBuilder metadatas.Builder,
	chunkBuilder chunks.Builder,
	chunkBasePath []string,
	minSizeToChunkInBytes uint,
) Application {
	out := application{
		hashAdapter:           hashAdapter,
		repository:            repository,
		service:               service,
		commitRepository:      commitRepository,
		fileRepository:        fileRepository,
		fileService:           fileService,
		databaseBuilder:       databaseBuilder,
		commitBuilder:         commitBuilder,
		executionsBuilder:     executionsBuilder,
		executionBuilder:      executionBuilder,
		metaDataBuilder:       metaDataBuilder,
		chunkBuilder:          chunkBuilder,
		chunkBasePath:         chunkBasePath,
		minSizeToChunkInBytes: minSizeToChunkInBytes,
		commits:               map[uint]commit{},
		contexts:              map[uint]contexts{},
	}

	return &out
}

// Begin begins a context on a database
func (app *application) Begin(path []string) (*uint, error) {
	if !app.repository.Exists(path) {
		str := fmt.Sprintf("the database (path: %s) does not currently exists and therefore must be initialized using the BeginWithInit method", filepath.Join(path...))
		return nil, errors.New(str)
	}

	return app.begin(path, "", "")
}

// BeginWithInit begins with init
func (app *application) BeginWithInit(path []string, name string, description string) (*uint, error) {
	if app.repository.Exists(path) {
		str := fmt.Sprintf("the database (path: %s) already exists and therefore must NOT be initialized, please use the Begin method directly", filepath.Join(path...))
		return nil, errors.New(str)
	}

	return app.begin(path, name, description)
}

func (app *application) begin(path []string, name string, description string) (*uint, error) {
	if !app.fileRepository.Exists(path) {
		err := app.fileService.Init(path)
		if err != nil {
			return nil, err
		}
	}

	err := app.fileService.Lock(path)
	if err != nil {
		return nil, err
	}

	contextStr := contexts{
		path:       path,
		executions: []executionData{},
	}

	if name != "" && description != "" {
		metaData, err := app.metaDataBuilder.Create().
			WithPath(path).
			WithName(name).
			WithDescription(description).
			Now()

		if err != nil {
			return nil, err
		}

		contextStr.metaData = metaData
	}

	keyname := uint(len(app.contexts))
	app.contexts[keyname] = contextStr

	return &keyname, nil
}

// Execute executes an execution on a context
func (app *application) Execute(context uint, bytes []byte) error {
	if contextIns, ok := app.contexts[context]; ok {
		sizeInBytes := uint(len(bytes))
		isChk := sizeInBytes >= app.minSizeToChunkInBytes
		newExecution := executionData{
			bytes: bytes,
		}

		builder := app.executionBuilder.Create()
		if isChk {
			pFinger, err := app.hashAdapter.FromBytes(bytes)
			if err != nil {
				return err
			}

			fingerStr := pFinger.String()
			split := app.splitString(fingerStr, splitHashInSubDirAmount)
			fullDir := append(app.chunkBasePath, split...)

			chk, err := app.chunkBuilder.Create().
				WithFingerPrint(*pFinger).
				WithPath(fullDir).
				Now()

			if err != nil {
				return err
			}

			builder.WithChunk(chk)
		}

		if !isChk {
			newExecution.bytes = nil
			builder.WithBytes(bytes)
		}

		execution, err := builder.Now()
		if err != nil {
			return err
		}

		newExecution.execution = execution
		contextIns.executions = append(contextIns.executions, newExecution)
		app.contexts[context] = contextIns
		return nil
	}

	str := fmt.Sprintf(invalidContextErrorPattern, context)
	return errors.New(str)
}

// Batch executes a batch executions on a context
func (app *application) Batch(context uint, bytes [][]byte) error {
	for idx, oneBytes := range bytes {
		err := app.Execute(context, oneBytes)
		if err != nil {
			str := fmt.Sprintf("there was an error while processing the data at index (%d) while executing the Batch: %s", idx, err.Error())
			return errors.New(str)
		}
	}

	return nil
}

// Commit executes a commit on a context
func (app *application) Commit(context uint) error {
	if contextIns, ok := app.contexts[context]; ok {
		executionsList := []executions.Execution{}
		for _, oneExecutionData := range contextIns.executions {
			executionsList = append(executionsList, oneExecutionData.execution)
			if oneExecutionData.execution.IsChunk() {
				err := app.fileService.Save(
					oneExecutionData.execution.Chunk().Path(),
					oneExecutionData.bytes,
				)

				if err != nil {
					return err
				}
			}
		}

		executions, err := app.executionsBuilder.Create().
			WithList(executionsList).
			Now()

		if err != nil {
			return err
		}

		commitBuilder := app.commitBuilder.Create().
			WithExecutions(executions)

		if app.repository.Exists(contextIns.path) {
			prevDatabase, err := app.repository.Retrieve(contextIns.path)
			if err != nil {
				return err
			}

			head := prevDatabase.Head()
			if head.HasParent() {
				parent := head.Parent()
				commitBuilder.WithParent(parent)
			}

		}

		commitIns, err := commitBuilder.Now()
		if err != nil {
			return err
		}

		commitsList := []commits.Commit{}
		if _, ok := app.commits[context]; ok {
			commitsList = app.commits[context].commits
		}

		commitsList = append(commitsList, commitIns)
		app.commits[context] = commit{
			path:     contextIns.path,
			commits:  commitsList,
			metaData: contextIns.metaData,
		}

		app.contexts[context] = contexts{
			path:       contextIns.path,
			executions: []executionData{},
			metaData:   nil,
		}

		return nil

	}

	str := fmt.Sprintf(invalidContextErrorPattern, context)
	return errors.New(str)
}

// Cancel executes a cancel on a context
func (app *application) Cancel(context uint) {
	delete(app.contexts, context)
	delete(app.commits, context)
}

// Push pushes updates of a context to its database
func (app *application) Push(context uint) error {
	if commitIns, ok := app.commits[context]; ok {
		if commitIns.metaData == nil {
			database, err := app.repository.Retrieve(commitIns.path)
			if err != nil {
				return err
			}

			commitIns.metaData = database.MetaData()
		}

		list := []databases.Database{}
		for _, oneCommit := range commitIns.commits {
			updatedDatabase, err := app.databaseBuilder.Create().
				WithMetaData(commitIns.metaData).
				WithHead(oneCommit).
				Now()

			if err != nil {
				return err
			}

			list = append(list, updatedDatabase)
		}

		err := app.service.SaveAll(list)
		if err != nil {
			return err
		}

		err = app.fileService.Unlock(commitIns.metaData.Path())
		if err != nil {
			return err
		}

		delete(app.commits, context)
		return nil
	}

	str := fmt.Sprintf(noCommitForContextErrorPattern, context)
	return errors.New(str)
}

// RollbackToPrevious rollback to the previous state
func (app *application) RollbackToPrevious(context uint) error {
	if contextIns, ok := app.contexts[context]; ok {
		database, err := app.repository.Retrieve(contextIns.path)
		if err != nil {
			return err
		}

		parent := database.Head().Parent()
		parentCommit, err := app.commitRepository.Retrieve(parent)
		if err != nil {
			return err
		}

		metaData := database.MetaData()
		updatedDatabase, err := app.databaseBuilder.Create().
			WithHead(parentCommit).
			WithMetaData(metaData).
			Now()

		if err != nil {
			return err
		}

		return app.service.Save(updatedDatabase)

	}

	str := fmt.Sprintf(invalidContextErrorPattern, context)
	return errors.New(str)
}

// RollbackToState rollback to the provided commmit state
func (app *application) RollbackToState(context uint, headCommit hash.Hash) error {
	if contextIns, ok := app.contexts[context]; ok {
		database, err := app.repository.Retrieve(contextIns.path)
		if err != nil {
			return err
		}

		parentCommit, err := app.commitRepository.Retrieve(headCommit)
		if err != nil {
			return err
		}

		metaData := database.MetaData()
		updatedDatabase, err := app.databaseBuilder.Create().
			WithHead(parentCommit).
			WithMetaData(metaData).
			Now()

		if err != nil {
			return err
		}

		return app.service.Save(updatedDatabase)

	}

	str := fmt.Sprintf(invalidContextErrorPattern, context)
	return errors.New(str)
}

func (app *application) splitString(str string, size int) []string {
	var parts []string
	partSize := len(str) / size
	for i := 0; i < size; i++ {
		start := i * partSize
		end := start + partSize
		if i == size-1 {
			end = len(str)
		}
		parts = append(parts, str[start:end])
	}

	return parts
}
