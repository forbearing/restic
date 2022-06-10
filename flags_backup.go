package restic

// - BackupFlags includes all flags of "restic backup" and inheris GlobalFlags
// - The "backup" command creates a new snapshot and saves the files and directories
type BackupFlags struct {
	// -n, --dry-run[=false]
	// do not upload or write any data, just show what would be done
	DryRun bool
	// -e, --exclude=[]
	// exclude a pattern (can be specified multiple times)
	Exclude []string
	// --exclude-caches[=false]
	// excludes cache directories that are marked with a CACHEDIR.TAG file.
	// See https://bford.info/cachedir/ for the Cache  Directory Tagging Standard
	ExcludeCache bool
	// --exclude-file=[]
	// read exclude patterns from a file (can be specified multiple times)
	ExcludeFile []string
	//--exclude-if-present=[]
	// takes filename[:header], exclude contents of directories containing
	// filename (except filename itself) if header of that file is as provided
	// (can be specified multiple times)
	ExcludeIfPresent []string
	// --exclude-larger-than=""
	// max size of the files to be backed up (allowed suffixes: k/K, m/M, g/G, t/T)
	ExcludeLargerThan string
	// --files-from=[]
	// read the files to backup from file (can be combined with file args;
	// can be specified multiple times)
	FilesFrom []string
	// --files-from-raw=[]
	// read the files to backup from file (can be combined with file args;
	// can be specified multiple times)
	FilesFromRaw []string
	// --files-from-verbatim=[]
	// read the files to backup from file (can be combined with file args;
	// can be specified multiple times)
	FilesFromVerbatim []string
	// -f, --force[=false]
	// force re-reading the target files/directories (overrides the "parent" flag)
	Force bool
	// -h, --help[=false]
	// help for backup
	Help bool
	// -H, --host=""
	// set the hostname for the snapshot manually. To prevent an expensive
	// rescan use the "parent" flag
	Host string
	// --iexclude=[]
	// same as --exclude pattern but ignores the casing of filenames
	Iexclude []string
	// --iexclude-file=[]
	// same as --exclude-file but ignores casing of filenames in patterns
	IexcludeFile []string
	// --ignore-ctime[=false]
	// ignore ctime changes when checking for modified files
	IgnoreCtime bool
	// --ignore-inode[=false]
	// ignore inode number changes when checking for modified files
	IgnoreInode bool
	// -x, --one-file-system[=false]
	// exclude other file systems, don't cross filesystem boundaries
	// and subvolumes
	OneFileSystem bool
	// --parent=""
	// use this parent snapshot (default: last snapshot in the repo
	// that has the same target files/directories, and is not newer
	// than the snapshot time)
	Parent string
	// --stdin[=false]
	// read backup from stdin
	Stdin bool
	// --stdin-filename="stdin"
	// filename to use when reading from stdin
	StdinFilename string
	// --tag=[]
	// add tags for the new snapshot in the format tag[,tag,...]
	// (can be specified multiple times)
	Tag []string
	// --time=""
	// time of the backup (ex. '2012-11-01 22:08:41') (default: now)
	Time string
	// --with-atime[=false]
	// store the atime for all files and directories
	WithAtime bool

	GlobalFlags
}

func concateBackupFlags() {
}
