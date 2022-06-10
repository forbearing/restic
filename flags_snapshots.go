package restic

// - SnapshotsFlags includes all flags of "restic snapshots" and inheris GlobalFlags
// - The "snapshots" command list all snapshots stored in the repository.
type SnapshotsFlags struct {
	// -c, --compact[=false]
	// use compact output format
	Compact bool
	// -g, --group-by=""
	// string for grouping snapshots by host,paths,tags
	GroupBy string
	// -h, --help[=false]
	// help for snapshots
	Help bool
	// -H, --host=[]
	// only consider snapshots for this host (can be specified multiple times)
	Host []string
	// --latest=0
	// only show the last n snapshots for each host and path
	Latest int
	// --path=[]
	// only consider snapshots for this path (can be specified multiple times)
	Path []string
	// --tag=[]
	// only consider snapshots which include this taglist in the format tag[,tag,...] (can be specified multiple times)
	Tag []string

	GlobalFlags
}

func concatSnapshotsFlags() {}
