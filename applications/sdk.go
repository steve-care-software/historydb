package applications

import (
	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/databases/commits/executions"
	"github.com/steve-care-software/historydb/domain/databases/commits/executions/chunks"
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/files"
	"github.com/steve-care-software/historydb/domain/hash"
)

const invalidContextErrorPattern = "the context, %d, is invalid"
const noCommitForContextErrorPattern = "there is no commit for the context %d"
const splitHashInSubDirAmount = 8

// NewApplication creates a new application
func NewApplication(
	repository databases.Repository,
	service databases.Service,
	commitRepository commits.Repository,
	fileRepository files.Repository,
	fileService files.Service,
	chunkBasePath []string,
	minSizeToChunkInBytes uint,
) Application {
	hashAdapter := hash.NewAdapter()
	databaseBuilder := databases.NewBuilder()
	commitBuilder := commits.NewBuilder()
	executionsBuilder := executions.NewBuilder()
	executionBuilder := executions.NewExecutionBuilder()
	metaDataBuilder := metadatas.NewBuilder()
	chunkBuilder := chunks.NewBuilder()
	return createApplication(
		hashAdapter,
		repository,
		service,
		commitRepository,
		fileRepository,
		fileService,
		databaseBuilder,
		commitBuilder,
		executionsBuilder,
		executionBuilder,
		metaDataBuilder,
		chunkBuilder,
		chunkBasePath,
		minSizeToChunkInBytes,
	)
}

// Application represents an application
type Application interface {
	Begin(path []string) (*uint, error)
	BeginWithInit(path []string, name string, description string) (*uint, error)
	Execute(context uint, bytes []byte) error
	Batch(context uint, bytes [][]byte) error
	Commit(context uint) error
	Cancel(context uint)
	Push(context uint) error
	RollbackToPrevious(context uint) error
	RollbackToState(context uint, headCommit hash.Hash) error
}
