package restic

// - RebuildIndexFlags includes all flags of "restic rebuild-index"
//   and inheris GlobalFlags.
// - The "rebuild-index" command creates a new index based on the pack files
//   in the repository.
type RebuildIndexFlags struct {
	// -h, --help[=false]
	// help for rebuild-index
	Help bool
	// --read-all-packs[=false]
	// read all pack files to generate new index from scratch
	ReadAllPacks bool

	GlobalFlags
}

func concatRebuildIndexFlags() {}
