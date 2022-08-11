package restic

import "strings"

// - Ls includes all flags of "restic ls" and inheris GlobalFlags
// - The "ls" command lists files and directories in a snapshots.
// - The special snapshot ID "latest" can be used to list files and directories
//   of the latest snapshot in the repository. The --host flag can be used in
//   conjunction to select the latest snapshot originating from a certain host only.
// - File listings can optionally be filtered by directories. Any positional
//   arguments after the snapshot ID are interpreted as absolute directory paths,
//   and only  files  inside  those  directories will be listed. If the --recursive
//   flag is used, then the filter will allow traversing into matching directories'
//   subfolders. Any directory paths specified must be absolute (starting with
//   a path separator); paths use the forward slash '/' as separator.
type Ls struct {
	// -h, --help[=false]
	// help for ls
	Help bool `json:"--help"`
	// -H, --host=[]
	// only consider snapshots for this host, when snapshot ID "latest" is given
	// (can be specified multiple times)
	Host []string `json:"--host"`
	// -l, --long[=false]
	// use a long listing format showing size and mode
	Long bool `json:"--long"`
	// --path=[]
	// only consider snapshots which include this (absolute) path, when
	// snapshot ID "latest" is given (can be specified multiple times)
	Path []string `json:"--path"`
	// --recursive[=false]
	// include files in subfolders of the listed directories
	Recursive bool `json:"--recursive"`
	// --tag=[]
	// only consider snapshots which include this taglist, when snapshot ID "latest"
	// is given (can be specified multiple times)
	Tag []string `json:"--tag"`

	args strings.Builder

	GlobalFlags
}

func (l *Ls) Flags() string {
	return concatFlags(l)
}

func (l *Ls) Name() string {
	return "ls"
}

func (l *Ls) Args() string {
	return strings.TrimSpace(l.args.String())
}

func (l *Ls) SetArgs(args ...string) string {
	for _, s := range args {
		l.args.WriteString(s + " ")
	}
	return l.Args()
}
