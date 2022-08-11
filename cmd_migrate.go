package restic

import "strings"

// Migrate includes all flags of "restic migrate" and inheris GlobalFlags.
// Migrate object implements the interface "Command".

// The "migrate" command applies migrations to a repository. When no migration
// name is explicitly given, a list of migrations that can be applies is printed.
type Migrate struct {
	// -f, --force[=false]
	// apply a migration a second time
	Force bool `json:"--force"`
	// -h, --help[=false]
	// help for migrate
	Help bool `json:"--help"`

	args strings.Builder

	GlobalFlags
}

func (m Migrate) Name() string  { return "migrate" }
func (m Migrate) Flags() string { return concatFlags(m) }
func (m Migrate) Args() string  { return strings.TrimSpace(m.args.String()) }

func (m Migrate) SetArgs(args ...string) *Migrate {
	for _, s := range args {
		m.args.WriteString(s + " ")
	}
	return &m
}
