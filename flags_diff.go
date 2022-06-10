package restic

// - DiffFlags includes all flags of "restic diff" and inheris GlobalFlags
// - The "diff" command show differences from the first to the second snapshot.
//   the first characters in each line display what was has happened to a
//   particular file or directory:
//       o +  The item was added
//       o -  The item was removed
//       o U  The metadata (access mode, timestamps, ...) for the item was updated
//       o M  The file's content was modified
//       o T  The type was changed, e.g. a file was made a symlink
type DiffFlags struct {
	// -h, --help[=false]
	// help for diff
	Help bool
	// --metadata[=false]
	// print changes in metadata
	Metadata bool

	GlobalFlags
}

func concatDiffFlags() {}
