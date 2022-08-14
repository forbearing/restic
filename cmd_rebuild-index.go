package restic

import "strings"

// RebuildIndex includes all flags of "restic rebuild-index" and inheris GlobalFlags.
// RebuildIndex object implements the interface "Command".
//
// The "rebuild-index" command creates a new index based on the pack files
// in the repository.
type RebuildIndex struct {
	// -h, --help[=false]
	// help for rebuild-index
	Help bool `json:"--help"`
	// --read-all-packs[=false]
	// read all pack files to generate new index from scratch
	ReadAllPacks bool `json:"--read-all-packs"`

	args strings.Builder

	GlobalFlags
}

func (r RebuildIndex) Name() string  { return "rebuild-index" }
func (r RebuildIndex) Flags() string { return concatFlags(r) }
func (r RebuildIndex) Args() string  { return strings.TrimSpace(r.args.String()) }

func (r RebuildIndex) SetArgs(args ...string) *RebuildIndex {
	for _, s := range args {
		r.args.WriteString(s + " ")
	}
	return &r
}
