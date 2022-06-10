package restic

// - TagFlags includes all flags of "restic tag" and inheris GlobalFlags.
// - The "tag" command allows you to modify tags on existing snapshots.
// - You can set/replace the entire set of tags on a snapshots, or add tags to/remove
///  tags from the existing set.
// - When no snapshot-ID is given. all snapshots matching the host, tag and path
//   filter criteria are modified.
type TagFlags struct {
	// --add=[]
	// tags which will be added to the existing tags in the format tag[,tag,...]
	// (can be given multiple times)
	Add []string
	// -h, --help[=false]      help for tag
	Help bool
	// -H, --host=[]
	// only consider snapshots for this host, when no snapshot ID is given
	// (can be specified multiple times)
	Host []string
	// --path=[]
	// only consider snapshots which include this (absolute) path,
	// when no snapshot-ID is given
	Path []string
	// --remove=[]
	// tags which will be removed from the existing tags in the format tag[,tag,...]
	// (can be given multiple times)
	Remove []string
	// --set=[]
	// tags which will replace the existing tags in the format tag[,tag,...]
	// (can be given multiple times)
	Set []string
	// --tag=[]
	// only consider snapshots which include this taglist, when no snapshot-ID is given
	Tag []string

	GlobalFlags
}

func concatTagFlags() {}
