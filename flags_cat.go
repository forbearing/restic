package restic

// - CatFlags includes all flags of "restic cat" and inheris GlobalFlags
// - The "cat" command is used to print internal objects to stdout
type CatFlags struct {
	// -h, --help[=false]
	// help for cat
	Help bool

	GlobalFlags
}

func conatCatFlags() {}
