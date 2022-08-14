package restic

import "strings"

// Diff includes all flags of "restic diff" and inheris GlobalFlags.
// Diff object implements the interface "Command".
//
// The "diff" command show differences from the first to the second snapshot.
// the first characters in each line display what was has happened to a
// particular file or directory:
//     o +  The item was added
//     o -  The item was removed
//     o U  The metadata (access mode, timestamps, ...) for the item was updated
//     o M  The file's content was modified
//     o T  The type was changed, e.g. a file was made a symlink
type Diff struct {
	// -h, --help[=false]
	// help for diff
	Help bool `json:"--help"`
	// --metadata[=false]
	// print changes in metadata
	Metadata bool `json:"--metadata"`

	args strings.Builder

	GlobalFlags
}

func (d Diff) Name() string  { return "diff" }
func (d Diff) Flags() string { return concatFlags(d) }
func (d Diff) Args() string  { return strings.TrimSpace(d.args.String()) }

func (d Diff) SetArgs(args ...string) *Diff {
	for _, s := range args {
		d.args.WriteString(s + " ")
	}
	return &d
}
