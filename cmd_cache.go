package restic

import "strings"

// - Cache includes all flags of "restic cache" and inheris GlobalFlags
// - The "cache" command allows listing and cleaning local cache directories
type Cache struct {
	// --cleanup[=false]
	// remove old cache directories
	Cleanup bool `json:"--cleanup"`
	// -h, --help[=false]
	// help for cache
	Help bool `json:"--help"`
	// --max-age=30
	// max age in days for cache directories to be considered old
	MaxAge int `json:"--max-age"`
	// --no-size[=false]
	// do not output the size of the cache directories
	NoSize bool `json:"--no-size"`

	args strings.Builder

	GlobalFlags
}

func (c *Cache) Flags() string {
	return concatFlags(c)
}

func (c *Cache) Name() string {
	return "cache"
}

func (c *Cache) Args() string {
	return strings.TrimSpace(c.args.String())
}

func (c *Cache) SetArgs(args ...string) string {
	for _, s := range args {
		c.args.WriteString(s + " ")
	}
	return c.Args()
}
