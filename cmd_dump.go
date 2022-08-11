package restic

import "strings"

// Dump includes all flags of "restic dump" and inheris GlobalFlags.
// Dump object implements the interface "Command".

// The "dump" command extracts files from a snapshot from the reposi. if a single
// file is selected, it printing its contents o stdout. Folder are output as a
// tar (default) or zip file containing the contents of the specified folder.
// Pass "/" as file name to dump the whole snapshot as an archive file.

// The special snapshot "latest" can be used to use the latest snapshot in the
// repository.
type Dump struct {
	// -a, --archive="tar"
	// set archive format as "tar" or "zip"
	Archive string `json:"--archive"`
	// -h, --help[=false]
	// help for dump
	Help bool `json:"--help"`
	// -H, --host=[]
	// only consider snapshots for this host when the snapshot ID is "latest"
	// (can be specified multiple times)
	Host []string `json:"--host"`
	// --path=[]
	// only consider snapshots which include this (absolute) path for snapshot ID "latest"
	Path []string `json:"--path"`
	// --tag=[]
	// only consider snapshots which include this taglist for snapshot ID "latest"
	Tag []string `json:"--tag"`

	args strings.Builder

	GlobalFlags
}

func (d Dump) Name() string  { return "dump" }
func (d Dump) Flags() string { return concatFlags(d) }
func (d Dump) Args() string  { return strings.TrimSpace(d.args.String()) }

func (d Dump) SetArgs(args ...string) *Dump {
	for _, s := range args {
		d.args.WriteString(s + " ")
	}
	return &d
}
