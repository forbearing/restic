package restic

import "strings"

// - Mount includes all flags of "restic mount" and inheris GlobalFlags
// - The "mount" command mounts the repository via fuse to a directory.
//   This is a read-only mount.
type Mount struct {
	// --allow-other[=false]
	// allow other users to access the data in the mounted directory
	AllowOther bool `json:"--allow-other"`
	// -h, --help[=false]
	// help for mount
	Help bool `json:"--help"`
	// -H, --host=[]
	// only consider snapshots for this host (can be specified multiple times)
	Host []string `json:"--host"`
	// --no-default-permissions[=false]
	// for 'allow-other', ignore Unix permissions and allow users to read all
	// snapshot files
	NoDefaultPermissions bool `json:"--no-default-permissions"`
	// --owner-root[=false]
	// use 'root' as the owner of files and dirs
	OwnerRoot bool `json:"--owner-root"`
	// --path=[]
	// only consider snapshots which include this (absolute) path
	Path []string `json:"--path"`
	// --snapshot-template="2006-01-02T15:04:05Z07:00"
	// set template to use for snapshot dirs
	SnapshotTemplate string `json:"--snapshot-template"`

	args strings.Builder

	GlobalFlags
}

func (m Mount) Name() string  { return "mount" }
func (m Mount) Flags() string { return concatFlags(m) }
func (m Mount) Args() string  { return strings.TrimSpace(m.args.String()) }

func (m Mount) SetArgs(args ...string) *Mount {
	for _, s := range args {
		m.args.WriteString(s + " ")
	}
	return &m
}
