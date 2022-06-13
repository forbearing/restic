package restic

import "strings"

// - Snapshots includes all flags of "restic snapshots" and inheris GlobalFlags
// - The "snapshots" command list all snapshots stored in the repository.
type Snapshots struct {
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

func (s *Snapshots) Flags() string {
	return concat(s)
}

func (s *Snapshots) Name() string {
	return "snapshots"
}

func (s *Snapshots) Args() string {
	return strings.TrimSpace(s.args.String())
}

func (s *Snapshots) SetArgs(args ...string) string {
	for _, str := range args {
		s.args.WriteString(str + " ")
	}
	return s.Args()
}
