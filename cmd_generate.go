package restic

import "strings"

// - Generate includes all flags of "restic generate" and inheris GlobalFlags
// - The "generate" command writes automatically generated files (like the man pages
//   and the auto-completion files for bash, fish and zsh).
type Generate struct {
	// --bash-completion=""
	// write bash completion file
	BashCompletion string `json:"--bash-completion"`
	// --fish-completion=""
	// write fish completion file
	FishCompletion string `json:"--fish-completion"`
	// -h, --help[=false]
	// help for generate
	Help bool `json:"--help"`
	// --man=""
	// write man pages to directory
	Man string `json:"--man"`
	// --zsh-completion=""
	// write zsh completion file
	ZshCompletion string `json:"--zsh-completion"`

	args strings.Builder

	GlobalFlags
}

func (g *Generate) Flags() string {
	return concatFlags(g)
}

func (g *Generate) Name() string {
	return "generate"
}

func (g *Generate) Args() string {
	return strings.TrimSpace(g.args.String())
}

func (g *Generate) SetArgs(args ...string) string {
	for _, s := range args {
		g.args.WriteString(s + " ")
	}
	return g.Args()
}
