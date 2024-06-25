package jsons

import (
	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/databases/commits/executions"
	"github.com/steve-care-software/historydb/domain/databases/commits/executions/chunks"
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewDatabaseAdapter creates a new database adapter
func NewDatabaseAdapter() databases.Adapter {
	metaDataBuilder := metadatas.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createDatabaseAdapter(
		metaDataBuilder,
		hashAdapter,
	)
}

// NewCommitAdapter creates a new commit adapter
func NewCommitAdapter() commits.Adapter {
	commitBuilder := commits.NewBuilder()
	executionsBuilder := executions.NewBuilder()
	executionBuilder := executions.NewExecutionBuilder()
	chunkBuilder := chunks.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createCommitAdapter(
		commitBuilder,
		executionsBuilder,
		executionBuilder,
		chunkBuilder,
		hashAdapter,
	)
}
