package restic

// - CacheFlags includes all flags of "restic cache" and inheris GlobalFlags
// - The "cache" command allows listing and cleaning local cache directories
type CacheFlags struct {
	// --cleanup[=false]
	// remove old cache directories
	Cleanup bool
	// -h, --help[=false]
	// help for cache
	Help bool
	// --max-age=30
	// max age in days for cache directories to be considered old
	MaxAge int
	// --no-size[=false]
	// do not output the size of the cache directories
	NoSize bool

	GlobalFlags
}

func concatCacheFlag() {}
