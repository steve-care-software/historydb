# HistoryDB
HistoryDB is a database system that records the tree of data recording in a database and provides options to  commit changes into states, revert to previous states and merge database trees.

## Create JSON application
```go
    basePath := []string{"databases", "my_database"}
	commitInnerPath := []string{"commits"}
	chunksInnerPath := []string{"chunks"}
	sizeToChunk := uint(1024)
	splitHashInThisAmount := uint(16)
	application := bundles.NewApplicationWithJSONAdapter(
		basePath,
		commitInnerPath,
		chunksInnerPath,
		sizeToChunk,
		splitHashInThisAmount,
	)
```