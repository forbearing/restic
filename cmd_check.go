package restic

import "strings"

// - Check includes all flags of "restic check" and inheris GlobalFlags
// - The "check" command tests the repository for errors and reports any errors it finds.
//   It can also be used to read all data and therefore simulate a restore.
// - By default, the "check" command will always load all data directly from the
//   repository and not use a local cache.
type Check struct {
	// --check-unused[=false]
	// find unused blobs
	CheckUnused bool `json:"--check-unused"`
	// -h, --help[=false]
	// help for check
	Help bool `json:"--help"`
	// --read-data[=false]
	// read all data blobs
	ReadData bool `json:"--read-data"`
	// --read-data-subset=""
	// read a subset of data packs, specified as 'n/t' for specific part,
	// or either 'x%' or 'x.y%' or a size in bytes with suffixes k/K, m/M,
	// g/G, t/T for a random subset
	ReadDataSubset string `json:"--read-data-subset"`
	// --with-cache[=false]
	// use the cache
	WithCache bool `json:"--with-cache"`

	args strings.Builder
	GlobalFlags
}

func (c *Check) Flags() string {
	return concat(c)
}

func (c *Check) Name() string {
	return "check"
}

func (c *Check) Args() string {
	return strings.TrimSpace(c.args.String())
}

func (c *Check) SetArgs(args ...string) string {
	for _, s := range args {
		c.args.WriteString(s + " ")
	}
	return c.Args()
}
