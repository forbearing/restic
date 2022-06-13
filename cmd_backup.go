package restic

import (
	"strings"
)

// - Backup includes all flags of "restic backup" and inheris GlobalFlags
// - The "backup" command creates a new snapshot and saves the files and directories
type Backup struct {
	// -n, --dry-run[=false]
	// do not upload or write any data, just show what would be done
	DryRun bool `json:"--dry-run"`
	// -e, --exclude=[]
	// exclude a pattern (can be specified multiple times)
	Exclude []string `json:"--exclude"`
	// --exclude-caches[=false]
	// excludes cache directories that are marked with a CACHEDIR.TAG file.
	// See https://bford.info/cachedir/ for the Cache  Directory Tagging Standard
	ExcludeCache bool `json:"--exclude-caches"`
	// --exclude-file=[]
	// read exclude patterns from a file (can be specified multiple times)
	ExcludeFile []string `json:"--exclude-file"`
	//--exclude-if-present=[]
	// takes filename[:header], exclude contents of directories containing
	// filename (except filename itself) if header of that file is as provided
	// (can be specified multiple times)
	ExcludeIfPresent []string `json:"--exclude-if-present"`
	// --exclude-larger-than=""
	// max size of the files to be backed up (allowed suffixes: k/K, m/M, g/G, t/T)
	ExcludeLargerThan string `json:"--exclude-larger-than"`
	// --files-from=[]
	// read the files to backup from file (can be combined with file args;
	// can be specified multiple times)
	FilesFrom []string `json:"--files-from"`
	// --files-from-raw=[]
	// read the files to backup from file (can be combined with file args;
	// can be specified multiple times)
	FilesFromRaw []string `json:"--files-from-raw"`
	// --files-from-verbatim=[]
	// read the files to backup from file (can be combined with file args;
	// can be specified multiple times)
	FilesFromVerbatim []string `json:"--files-from-verbatim"`
	// -b, --force[=false]
	// force re-reading the target files/directories (overrides the "parent" flag)
	Force bool `json:"--force"`
	// -h, --help[=false]
	// help for backup
	Help bool `json:"--help"`
	// -H, --host=""
	// set the hostname for the snapshot manually. To prevent an expensive
	// rescan use the "parent" flag
	Host string `json:"--host"`
	// --iexclude=[]
	// same as --exclude pattern but ignores the casing of filenames
	Iexclude []string `json:"--iexclude"`
	// --iexclude-file=[]
	// same as --exclude-file but ignores casing of filenames in patterns
	IexcludeFile []string `json:"--iexclude-file"`
	// --ignore-ctime[=false]
	// ignore ctime changes when checking for modified files
	IgnoreCtime bool `json:"ignore-ctime"`
	// --ignore-inode[=false]
	// ignore inode number changes when checking for modified files
	IgnoreInode bool `json:"--ignore-inode"`
	// -x, --one-file-system[=false]
	// exclude other file systems, don't cross filesystem boundaries
	// and subvolumes
	OneFileSystem bool `json:"--one-file-system"`
	// --parent=""
	// use this parent snapshot (default: last snapshot in the repo
	// that has the same target files/directories, and is not newer
	// than the snapshot time)
	Parent string `json:"--parent"`
	// --stdin[=false]
	// read backup from stdin
	Stdin bool `json:"--stdin"`
	// --stdin-filename="stdin"
	// filename to use when reading from stdin
	StdinFilename string `json:"--stdin-filename"`
	// --tag=[]
	// add tags for the new snapshot in the format tag[,tag,...]
	// (can be specified multiple times)
	Tag []string `json:"--tag"`
	// --time=""
	// time of the backup (ex. '2012-11-01 22:08:41') (default: now)
	Time string `json:"--time"`
	// --with-atime[=false]
	// store the atime for all files and directories
	WithAtime bool `json:"--with-atime"`

	args strings.Builder

	GlobalFlags
}

func (b *Backup) Flags() string {
	return concat(b)
}

func (b *Backup) Name() string {
	return "backup"
}

func (b *Backup) Args() string {
	return strings.TrimSpace(b.args.String())
}

func (b *Backup) SetArgs(args ...string) string {
	for _, s := range args {
		b.args.WriteString(s + " ")
	}
	return b.Args()
}
