package restic

import "strings"

// - Tag includes all flags of "restic tag" and inheris GlobalFlags.
// - The "tag" command allows you to modify tags on existing snapshots.
// - You can set/replace the entire set of tags on a snapshots, or add tags to/remove
///  tags from the existing set.
// - When no snapshot-ID is given. all snapshots matching the host, tag and path
//   filter criteria are modified.
type Tag struct {
	// --add=[]
	// tags which will be added to the existing tags in the format tag[,tag,...]
	// (can be given multiple times)
	Add []string `json:"--add"`
	// -h, --help[=false]      help for tag
	Help bool `json:"--help"`
	// -H, --host=[]
	// only consider snapshots for this host, when no snapshot ID is given
	// (can be specified multiple times)
	Host []string `json:"--host"`
	// --path=[]
	// only consider snapshots which include this (absolute) path,
	// when no snapshot-ID is given
	Path []string `json:"--path"`
	// --remove=[]
	// tags which will be removed from the existing tags in the format tag[,tag,...]
	// (can be given multiple times)
	Remove []string `json:"--remove"`
	// --set=[]
	// tags which will replace the existing tags in the format tag[,tag,...]
	// (can be given multiple times)
	Set []string `json:"--set"`
	// --tag=[]
	// only consider snapshots which include this taglist, when no snapshot-ID is given
	Tag []string `json:"--tag"`

	args strings.Builder

	GlobalFlags
}

func (t *Tag) Flags() string {
	return concat(t)
}

func (t *Tag) Name() string {
	return "tag"
}

func (t *Tag) Args() string {
	return strings.TrimSpace(t.args.String())
}

func (t *Tag) SetArgs(args ...string) string {
	for _, s := range args {
		t.args.WriteString(s)
	}
	return t.Args()
}
