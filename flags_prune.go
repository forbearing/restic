package restic

// - PruneFlags includes all flags of "restic prune" and inheris GlobalFlags.
// - The "prune" command checks the repository and removes data that is not
//   referenced and therefore not needed any more.
type PruneFlags struct {
	// -n, --dry-run[=false]
	// do not modify the repository, just print what would be done
	DryRun bool
	// -h, --help[=false]
	// help for prune
	Help bool
	// --max-repack-size=""
	// maximum size to repack (allowed suffixes: k/K, m/M, g/G, t/T)
	MaxRepackSize string
	// --max-unused="5%"
	// tolerate  given  limit of unused data (absolute value in bytes with
	// suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited')
	MaxUnused string
	// --repack-cacheable-only[=false]
	// only repack packs which are cacheable
	RepackCacheableOnly bool

	GlobalFlags
}

func concatPruneFlags() {}
