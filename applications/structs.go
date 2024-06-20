package applications

import (
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/databases/commits/executions"
)

type contexts struct {
	path       []string
	executions []executionData
}

type executionData struct {
	execution executions.Execution
	bytes     []byte
}

type commit struct {
	path    []string
	commits []commits.Commit
}
