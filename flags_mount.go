package restic

// - MountFlags includes all flags of "restic mount" and inheris GlobalFlags
// - The "mount" command mounts the repository via fuse to a directory.
//   This is a read-only mount.
type MountFlags struct {
	// --allow-other[=false]
	// allow other users to access the data in the mounted directory
	AllowOther bool
	// -h, --help[=false]
	// help for mount
	Help bool
	// -H, --host=[]
	// only consider snapshots for this host (can be specified multiple times)
	Host []string
	// --no-default-permissions[=false]
	// for 'allow-other', ignore Unix permissions and allow users to read all
	// snapshot files
	NoDefaultPermissions bool
	// --owner-root[=false]
	// use 'root' as the owner of files and dirs
	OwnerRoot bool
	// --path=[]
	// only consider snapshots which include this (absolute) path
	Path []string
	// --snapshot-template="2006-01-02T15:04:05Z07:00"
	// set template to use for snapshot dirs
	SnapshotTemplate string

	GlobalFlags
}

func concatMountFlags() {}
