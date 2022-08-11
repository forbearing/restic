package restic

import (
	"strings"
)

// - Find includes all flags of "restic find" and inheris GlobalFlags
// - The "find" command searches for files and directories in snapshots stored in
//   the repo. It can also be used to search for blobs or trees for troubleshooting.
type Find struct {
	// --blob[=false]
	// pattern is a blob-ID
	Blob bool `json:"--blob"`
	// -h, --help[=false]
	// help for find
	Help bool `json:"--help"`
	// -H, --host=[]
	// only consider snapshots for this host, when no snapshot ID is given
	// (can be specified multiple times)
	Host []string `json:"--host"`
	// -i, --ignore-case[=false]
	// ignore case for pattern
	IgnoreCase bool `json:"--ignore-case"`
	// -l, --long[=false]
	// use a long listing format showing size and mode
	Long bool `json:"--long"`
	//-N, --newest=""      newest modification date/time
	Newest string `json:"--newest"`
	//-O, --oldest=""      oldest modification date/time
	Oldest string `json:"--oldest"`
	// --pack[=false]
	// pattern is a pack-ID
	Pack bool `json:"--pack"`
	// --path=[]
	// only consider snapshots which include this (absolute) path,
	// when no snapshot-ID is given
	Path []string `json:"--path"`
	// --show-pack-id[=false]
	// display the pack-ID the blobs belong to (with --blob or --tree)
	ShowPackId bool `json:"--show-pack-id"`
	// -s, --snapshot=[]
	// snapshot id to search in (can be given multiple times)
	Snapshot []string `json:"--snapshot"`
	// --tag=[]
	// only consider snapshots which include this taglist,
	// when no snapshot-ID is given
	Tag []string `json:"--tag"`
	// --tree[=false]
	// pattern is a tree-ID
	Tree bool `json:"--tree"`

	args strings.Builder

	GlobalFlags
}

func (f *Find) Flags() string {
	return concatFlags(f)
}

func (f *Find) Name() string {
	return "find"
}

func (f *Find) Args() string {
	return strings.TrimSpace(f.args.String())
}

func (f *Find) SetArgs(args ...string) string {
	for _, s := range args {
		f.args.WriteString(s + " ")
	}
	return f.Args()
}
