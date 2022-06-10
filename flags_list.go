package restic

// - ListFlags includes all flags of "restic list" and inheris GlobalFlags
// - The "list" command allows listing objects in the repository based on type
// - Usage restic list [flags] [blobs|packs|index|snapshots|keys|locks]
type ListFlags struct {
	// -h, --help[=false]
	// help for list
	Help bool

	GlobalFlags
}

func concatListFlags() {}
