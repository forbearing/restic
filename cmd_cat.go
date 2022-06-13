package restic

import "strings"

// - Cat includes all flags of "restic cat" and inheris GlobalFlags
// - The "cat" command is used to print internal objects to stdout
type Cat struct {
	// -h, --help[=false]
	// help for cat
	Help bool `json:"--help"`

	args strings.Builder

	GlobalFlags
}

func (c *Cat) Flags() string {
	return concat(c)
}

func (c *Cat) Name() string {
	return "cat"
}

func (c *Cat) Args() string {
	return strings.TrimSpace(c.args.String())
}

func (c *Cat) SetArgs(args ...string) string {
	for _, s := range args {
		c.args.WriteString(s + " ")
	}
	return c.Args()
}
