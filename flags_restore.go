package restic

// - RestoreFlags includes all flags of "restic restore" and inheris GlobalFlags
// - The "restore" command extracts the data from a snapshot from the repository
//   to a directory.
// - The special snapshot "latest" can be used to restore the latest snapshot
//   in the repository.
type RestoreFlags struct {
	// -e, --exclude=[]
	// exclude a pattern (can be specified multiple times)
	Exclude []string
	// -h, --help[=false]
	// help for restore
	Help bool
	// -H, --host=[]
	// only consider snapshots for this host when the snapshot ID is "latest"
	// (can be specified multiple times)
	Host []string
	// --iexclude=[]
	// same as --exclude but ignores the casing of filenames
	Iexclude []string
	// --iinclude=[]
	// same as --include but ignores the casing of filenames
	Iinclude []string
	// -i, --include=[]
	// include a pattern, exclude everything else (can be specified multiple times)
	Include []string
	// --path=[]
	// only consider snapshots which include this (absolute) path for snapshot ID "latest"
	Path []string
	// --tag=[]
	// only consider snapshots which include this taglist for snapshot ID "latest"
	Tag []string
	// -t, --target=""
	// directory to extract data to
	Target string
	// --verify[=false]
	// verify restored files content
	Verify bool

	GlobalFlags
}

func concatRestoreFlags() {}
