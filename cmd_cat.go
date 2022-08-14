package restic

import "strings"

// Cat includes all flags of "restic cat" and inheris GlobalFlags.
// Cat object implements the interface "Command".
//
// The "cat" command is used to print internal objects to stdout.
type Cat struct {
	// -h, --help[=false]
	// help for cat
	Help bool `json:"--help"`

	args strings.Builder

	GlobalFlags
}

func (c Cat) Name() string  { return "cat" }
func (c Cat) Flags() string { return concatFlags(c) }
func (c Cat) Args() string  { return strings.TrimSpace(c.args.String()) }

func (c Cat) SetArgs(args ...string) *Cat {
	for _, s := range args {
		c.args.WriteString(s + " ")
	}
	return &c
}
