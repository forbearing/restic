package restic

import "strings"

// - SnapshotsFlags includes all flags of "restic snapshots" and inheris GlobalFlags
// - The "snapshots" command list all snapshots stored in the repository.
type SnapshotsFlags struct {
	// -c, --compact[=false]
	// use compact output format
	Compact bool `json:"--compact"`
	// -g, --group-by=""
	// string for grouping snapshots by host,paths,tags
	GroupBy string `json:"--group-by"`
	// -h, --help[=false]
	// help for snapshots
	Help bool `json:"--help"`
	// -H, --host=[]
	// only consider snapshots for this host (can be specified multiple times)
	Host []string `json:"--host"`
	// --latest=0
	// only show the last n snapshots for each host and path
	Latest int `json:"--latest"`
	// --path=[]
	// only consider snapshots for this path (can be specified multiple times)
	Path []string `json:"--path"`
	// --tag=[]
	// only consider snapshots which include this taglist in the format tag[,tag,...] (can be specified multiple times)
	Tag []string `json:"--tag"`

	args strings.Builder

	GlobalFlags
}

func (f *SnapshotsFlags) Flags() string {
	return concat(f)
}

func (f *SnapshotsFlags) Name() string {
	return "snapshots"
}

func (f *SnapshotsFlags) Args() string {
	return f.args.String()
}

func (f *SnapshotsFlags) SetArgs(args ...string) string {
	for _, s := range args {
		f.args.WriteString(s)
	}
	return f.Args()
}
