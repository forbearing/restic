package restic

import "strings"

// - Forget includes all flags of "restic forget" and inheris GlobalFlags
// - The "forget" command removes snapshots according to a policy. please note
//   that this command really only deletes the snapshot object in the repository,
//   which is a reference to data stored there. In order to remove the unreferenced
//   data after "forget" was run successfully, see the "prune" command. Please
//   alse read the documentation for "forget" to lean about important security considerations.
type Forget struct {
	// -l, --keep-last=0
	// keep the last n snapshots
	KeepLast int `json:"--keep-last"`
	// -H, --keep-hourly=0
	// keep the last n hourly snapshots
	KeepHourly int `json:"--keep-hourly"`
	// -d, --keep-daily=0
	// keep the last n daily snapshots
	KeepDaily int `json:"--keep-daily"`
	// -w, --keep-weekly=0
	// keep the last n weekly snapshots
	KeepWeekly int `json:"--keep-weekly"`
	// -m, --keep-monthly=0
	// keep the last n monthly snapshots
	KeepMonthly int `json:"--keep-monthly"`
	// -y, --keep-yearly=0
	// keep the last n yearly snapshots
	KeepYearly int `json:"--keep-yearly"`
	// --keep-within=
	// keep snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithin string `json:"--keep-within"`
	// --keep-within-hourly=
	// keep hourly snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinHourly string `json:"--keep-within-hourly"`
	// --keep-within-daily=
	// keep daily snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinDaily string `json:"--keep-within-daily"`
	// --keep-within-weekly=
	// keep weekly snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinWeekly string `json:"--keep-within-weekly"`
	// --keep-within-monthly=
	// keep monthly snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinMonthly string `json:"--keep-within-monthly"`
	// --keep-within-yearly=
	// keep yearly snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinYearly string `json:"--keep-within-yearly"`
	// --keep-tag=[]
	// keep snapshots with this taglist (can be specified multiple times)
	KeepTag []string `json:"--keep-tag"`
	// --host=[]
	// only consider snapshots with the given host (can be specified multiple times)
	Host []string `json:"--host"`
	// --tag=[]
	// only consider snapshots which include this taglist in the format tag[,tag,...]
	// (can be specified multiple times)
	Tag []string `json:"--tag"`
	// --path=[]
	// only consider snapshots which include this (absolute) path
	// (can be specified multiple times)
	Path []string `json:"--path"`
	// -c, --compact[=false]
	// use compact output format
	Compact bool `json:"--compact"`
	// -g, --group-by="host,paths"
	// string for grouping snapshots by host,paths,tags
	GroupBy string `json:"--group-by"`
	// -n, --dry-run[=false]
	// do not delete anything, just print what would be done
	DryRun bool `json:"--dry-run"`
	// --prune[=false]
	// automatically run the 'prune' command if snapshots have been removed
	Prune bool `json:"--prune"`
	// --max-unused="5%"
	// tolerate  given  limit of unused data (absolute value in bytes with
	// suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited')
	MaxUnused string `json:"--max-unused"`
	// --max-repack-size=""
	// maximum size to repack (allowed suffixes: k/K, m/M, g/G, t/T)
	MaxRepackSize string `json:"--max-repack-size"`
	// --repack-cacheable-only[=false]
	// only repack packs which are cacheable
	RepackCacheableOnly bool `json:"--repack-cacheable-only"`
	// -h, --help[=false]
	// help for forget
	Help bool `json:"--help"`

	args strings.Builder

	GlobalFlags
}

func (f *Forget) Flags() string {
	return concat(f)
}

func (f *Forget) Name() string {
	return "forget"
}

func (f *Forget) Args() string {
	return strings.TrimSpace(f.args.String())
}

func (f *Forget) SetArgs(args ...string) string {
	for _, s := range args {
		f.args.WriteString(s + " ")
	}
	return f.Args()
}
