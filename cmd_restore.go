package restic

import "strings"

// - Restore includes all flags of "restic restore" and inheris GlobalFlags
// - The "restore" command extracts the data from a snapshot from the repository
//   to a directory.
// - The special snapshot "latest" can be used to restore the latest snapshot
//   in the repository.
type Restore struct {
	// -e, --exclude=[]
	// exclude a pattern (can be specified multiple times)
	Exclude []string `json:"--exclude"`
	// -h, --help[=false]
	// help for restore
	Help bool `json:"--help"`
	// -H, --host=[]
	// only consider snapshots for this host when the snapshot ID is "latest"
	// (can be specified multiple times)
	Host []string `json:"--host"`
	// --iexclude=[]
	// same as --exclude but ignores the casing of filenames
	Iexclude []string `json:"--iexclude"`
	// --iinclude=[]
	// same as --include but ignores the casing of filenames
	Iinclude []string `json:"--iinclude"`
	// -i, --include=[]
	// include a pattern, exclude everything else (can be specified multiple times)
	Include []string `json:"--include"`
	// --path=[]
	// only consider snapshots which include this (absolute) path for snapshot ID "latest"
	Path []string `json:"--path"`
	// --tag=[]
	// only consider snapshots which include this taglist for snapshot ID "latest"
	Tag []string `json:"--tag"`
	// -t, --target=""
	// directory to extract data to
	Target string `json:"--target"`
	// --verify[=false]
	// verify restored files content
	Verify bool `json:"--verify"`

	args strings.Builder

	GlobalFlags
}

func (r *Restore) Flags() string {
	return concat(r)
}

func (r *Restore) Name() string {
	return "restore"
}

func (r *Restore) Args() string {
	return strings.TrimSpace(r.args.String())
}

func (r *Restore) SetArgs(args ...string) string {
	for _, s := range args {
		r.args.WriteString(s + " ")
	}
	return r.Args()
}
