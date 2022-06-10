package restic

// - CheckFlags includes all flags of "restic check" and inheris GlobalFlags
// - The "check" command tests the repository for errors and reports any errors it finds.
//   It can also be used to read all data and therefore simulate a restore.
// - By default, the "check" command will always load all data directly from the
//   repository and not use a local cache.
type CheckFlags struct {
	// --check-unused[=false]
	// find unused blobs
	CheckUnused bool
	// -h, --help[=false]
	// help for check
	Help bool
	// --read-data[=false]
	// read all data blobs
	ReadData bool
	// --read-data-subset=""
	// read a subset of data packs, specified as 'n/t' for specific part,
	// or either 'x%' or 'x.y%' or a size in bytes with suffixes k/K, m/M,
	// g/G, t/T for a random subset
	ReadDataSubset string
	// --with-cache[=false]
	// use the cache
	WithCache bool

	GlobalFlags
}

func concatCheckFlags() {}
