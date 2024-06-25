package applications

import (
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/databases/commits/executions"
	"github.com/steve-care-software/historydb/domain/databases/metadatas"
)

type contexts struct {
	path       []string
	executions []executionData
	metaData   metadatas.MetaData
}

type executionData struct {
	execution executions.Execution
	bytes     []byte
}

type commit struct {
	path     []string
	commits  []commits.Commit
	metaData metadatas.MetaData
}
