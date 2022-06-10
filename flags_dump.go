package restic

// - DumpFlags includes all flags of "restic dump" and inheris GlobalFlags
// - The "dump" command extracts files from a snapshot from the reposi. if a single
//   file is selected, it printing its contents o stdout. Folder are output as a
//   tar (default) or zip file containing the contents of the specified folder.
//   Pass "/" as file name to dump the whole snapshot as an archive file.
// - The special snapshot "latest" can be used to use the latest snapshot in the
//   repository.
type DumpFlags struct {
	// -a, --archive="tar"
	// set archive format as "tar" or "zip"
	Archive string
	// -h, --help[=false]
	// help for dump
	Help bool
	// -H, --host=[]
	// only consider snapshots for this host when the snapshot ID is "latest"
	// (can be specified multiple times)
	Host []string
	// --path=[]
	// only consider snapshots which include this (absolute) path for snapshot ID "latest"
	Path []string
	// --tag=[]
	// only consider snapshots which include this taglist for snapshot ID "latest"
	Tag []string

	GlobalFlags
}

func concatDumpFlags() {}
