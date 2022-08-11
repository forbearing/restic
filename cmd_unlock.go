package restic

import "strings"

// - Unlock includes all flags of "retic unlock" and inheris GlobalFlags
// - the "unlock" command removes stale locks that have been created by other
//   restic processes.
type Unlock struct {
	// -h, --help[=false]
	// help for unlock
	Help bool `json:"--help"`
	// --remove-all[=false]
	// remove all locks, even non-stale ones
	RemoveAll bool `json:"--remove-all"`

	args strings.Builder

	GlobalFlags
}

func (u *Unlock) Flags() string {
	return concatFlags(u)
}

func (u *Unlock) Name() string {
	return "unlock"
}

func (u *Unlock) Args() string {
	return strings.TrimSpace(u.args.String())
}
func (u *Unlock) SetArgs(args ...string) string {
	for _, s := range args {
		u.args.WriteString(s + " ")
	}
	return u.Args()
}
