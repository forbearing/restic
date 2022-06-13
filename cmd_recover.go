package restic

import "strings"

// - Recover includes all flags of "restic recover" and inheris GlobalFlags
// - The "recover" command builds a new repository from all directories it can find
//   in the raw data of the repository which are not referenced in an existing
//   snapshot. It can used if, for example, a snapshot has been removed by accident
//   with "forget"
type Recover struct {
	// -h, --help[=false]      help for recover
	Help bool `json:"--help"`

	args strings.Builder

	GlobalFlags
}

func (r *Recover) Flags() string {
	return concat(r)
}

func (r *Recover) Name() string {
	return "recover"
}

func (r *Recover) Args() string {
	return strings.TrimSpace(r.args.String())
}

func (r *Recover) SetArgs(args ...string) string {
	for _, s := range args {
		r.args.WriteString(s + " ")
	}
	return r.Args()
}
