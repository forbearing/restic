package restic

// - FindOption includes all flags of "restic find" and inheris GlobalFlags
// - The "find" command searches for files and directories in snapshots stored in
//   the repo. It can also be used to search for blobs or trees for troubleshooting.
type FindFlags struct {
	// --blob[=false]
	// pattern is a blob-ID
	Blob bool
	// -h, --help[=false]
	// help for find
	Help bool
	// -H, --host=[]
	// only consider snapshots for this host, when no snapshot ID is given
	// (can be specified multiple times)
	Host []string
	// -i, --ignore-case[=false]
	// ignore case for pattern
	IgnoreCase bool
	// -l, --long[=false]
	// use a long listing format showing size and mode
	Long bool
	//-N, --newest=""      newest modification date/time
	Newest string
	//-O, --oldest=""      oldest modification date/time
	Oldest string
	// --pack[=false]
	// pattern is a pack-ID
	Pack bool
	// --path=[]
	// only consider snapshots which include this (absolute) path,
	// when no snapshot-ID is given
	Path []string
	// --show-pack-id[=false]
	// display the pack-ID the blobs belong to (with --blob or --tree)
	ShowPackId bool
	// -s, --snapshot=[]
	// snapshot id to search in (can be given multiple times)
	Snapshot []string
	// --tag=[]
	// only consider snapshots which include this taglist,
	// when no snapshot-ID is given
	Tag []string
	// --tree[=false]
	// pattern is a tree-ID
	Tree bool

	GlobalFlags
}

func concatFindFlags() {}
