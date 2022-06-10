package restic

// - StatsFlags includes all flags of "restic stats" and inheris GlobalFlags.
// - The "stats" command walks one or multiple snapshots in a repository and
//   accumulates statistics about the data stored therein. It reports on the
//   number of unique files and their sizes, according to one of the counting
//   modes as given by the --mode flag.
// - It Operates on all snapshots matching the selection criteria or all snapshots
//   if northing is specified. the special snapshot ID "latest" is also supported.
//   some modes make more sense over just a single snapshot, while others are
//   useful across all snapshots, depending on what you are trying to calculate.
// - The mods are:
//       o restore-size: (default) Counts the size of the restored files.
//       o files-by-contents: Counts total size of files, where a file is
//         considered unique if it has unique contents.
//       o raw-data: Counts the size of blobs in the repository, regardless of
//         how many files reference them.
//       o blobs-per-file: A combination of files-by-contents and raw-data.
// - Refer to the online manual for more details about each mode.
type StatsFlags struct {
	// -h, --help[=false]
	// help for stats
	Help bool
	// -H, --host=[]
	// only consider snapshots with the given host (can be specified multiple times)
	Host []string
	// --mode="restore-size"
	// counting mode: restore-size (default), files-by-contents,
	//  blobs-per-file or raw-data
	Mode string
	// --path=[]
	// only consider snapshots which include this (absolute) path
	// (can be specified multiple times)
	Path []string
	// --tag=[]
	// only consider snapshots which include this taglist in the format tag[,tag,...]
	// (can be specified multiple times)
	Tag []string

	GlobalFlags
}

func concatStatsFlags() {}
