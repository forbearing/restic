package restic

// - UnlockFlags includes all flags of "retic unlock" and inheris GlobalFlags
// - the "unlock" command removes stale locks that have been created by other
//   restic processes.
type UnlockFlags struct {
	// -h, --help[=false]
	// help for unlock
	Help bool
	// --remove-all[=false]
	// remove all locks, even non-stale ones
	RemoveAll bool

	GlobalFlags
}

func concatUnlockFlags() {}
