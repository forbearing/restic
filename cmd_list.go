package restic

import "strings"

// - List includes all flags of "restic list" and inheris GlobalFlags
// - The "list" command allows listing objects in the repository based on type
// - Usage restic list [flags] [blobs|packs|index|snapshots|keys|locks]
type List struct {
	// -h, --help[=false]
	// help for list
	Help bool `json:"--help"`

	args strings.Builder

	GlobalFlags
}

func (l List) Name() string  { return "list" }
func (l List) Flags() string { return concatFlags(l) }
func (l List) Args() string  { return strings.TrimSpace(l.args.String()) }

func (l List) SetArgs(args ...string) *List {
	for _, s := range args {
		l.args.WriteString(s + " ")
	}
	return &l
}
