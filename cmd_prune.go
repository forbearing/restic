package restic

import "strings"

// - Prune includes all flags of "restic prune" and inheris GlobalFlags.
// - The "prune" command checks the repository and removes data that is not
//   referenced and therefore not needed any more.
type Prune struct {
	// -n, --dry-run[=false]
	// do not modify the repository, just print what would be done
	DryRun bool `json:"--dry-run"`
	// -h, --help[=false]
	// help for prune
	Help bool `json:"--help"`
	// --max-repack-size=""
	// maximum size to repack (allowed suffixes: k/K, m/M, g/G, t/T)
	MaxRepackSize string `json:"--max-repack-size"`
	// --max-unused="5%"
	// tolerate  given  limit of unused data (absolute value in bytes with
	// suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited')
	MaxUnused string `json:"--max-unused"`
	// --repack-cacheable-only[=false]
	// only repack packs which are cacheable
	RepackCacheableOnly bool `json:"--repack-cacheable-only"`

	args strings.Builder

	GlobalFlags
}

func (p *Prune) Flags() string {
	return concatFlags(p)
}

func (p *Prune) Name() string {
	return "prune"
}

func (p *Prune) Args() string {
	return strings.TrimSpace(p.args.String())
}

func (p *Prune) SetArgs(args ...string) string {
	for _, s := range args {
		p.args.WriteString(s + " ")
	}
	return p.Args()
}
