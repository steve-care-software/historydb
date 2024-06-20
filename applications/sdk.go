package applications

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

const invalidContextErrorPattern = "the context, %d, is invalid"
const noCommitForContextErrorPattern = "there is no commit for the context %d"
const splitHashInSubDirAmount = 8

// Application represents an application
type Application interface {
	Begin(path []string) (*uint, error)
	Execute(context uint, bytes []byte) error
	Batch(context uint, bytes [][]byte) error
	Commit(context uint) error
	Cancel(context uint)
	Push(context uint) error
	RollbackToPrevious(context uint) error
	RollbackToState(context uint, headCommit hash.Hash) error
}
