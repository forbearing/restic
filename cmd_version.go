package restic

import "strings"

// - Version includes all flags of "restic version" and inheris GlobalFlags.
// - The "version" command prints detailed information about the restic environment
//   and the version of the software
type Version struct {
	// -h, --help[=false]
	// help for version
	Help bool `json:"--help"`

	args strings.Builder

	GlobalFlags
}

func (v Version) Name() string  { return "version" }
func (v Version) Flags() string { return concatFlags(v) }
func (v Version) Args() string  { return strings.TrimSpace(v.args.String()) }

func (v Version) SetArgs(args ...string) *Version {
	for _, s := range args {
		v.args.WriteString(s + " ")
	}
	return &v
}
