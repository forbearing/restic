package restic

// - ForgetOption includes all flags of "restic forget" and inheris GlobalFlags
// - The "forget" command removes snapshots according to a policy. please note
//   that this command really only deletes the snapshot object in the repository,
//   which is a reference to data stored there. In order to remove the unreferenced
//   data after "forget" was run successfully, see the "prune" command. Please
//   alse read the documentation for "forget" to lean about important security considerations.
type ForgetFlags struct {
	// -l, --keep-last=0
	// keep the last n snapshots
	KeepLast int
	// -H, --keep-hourly=0
	// keep the last n hourly snapshots
	KeepHourly int
	// -d, --keep-daily=0
	// keep the last n daily snapshots
	KeepDaily int
	// -w, --keep-weekly=0
	// keep the last n weekly snapshots
	KeepWeekly int
	// -m, --keep-monthly=0
	// keep the last n monthly snapshots
	KeepMonthly int
	// -y, --keep-yearly=0
	// keep the last n yearly snapshots
	KeepYearly int
	// --keep-within=
	// keep snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithin string
	// --keep-within-hourly=
	// keep hourly snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinHourly string
	// --keep-within-daily=
	// keep daily snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinDaily string
	// --keep-within-weekly=
	// keep weekly snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinWeekly string
	// --keep-within-monthly=
	// keep monthly snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinMonthly string
	// --keep-within-yearly=
	// keep yearly snapshots that are newer than duration (eg. 1y5m7d2h)
	// relative to the latest snapshot
	KeepWithinYearly string
	// --keep-tag=[]
	// keep snapshots with this taglist (can be specified multiple times)
	KeepTag []string
	// --host=[]
	// only consider snapshots with the given host (can be specified multiple times)
	Host []string
	// --tag=[]
	// only consider snapshots which include this taglist in the format tag[,tag,...]
	// (can be specified multiple times)
	Tag []string
	// --path=[]
	// only consider snapshots which include this (absolute) path
	// (can be specified multiple times)
	Path []string
	// -c, --compact[=false]
	// use compact output format
	Compact bool
	// -g, --group-by="host,paths"
	// string for grouping snapshots by host,paths,tags
	GroupBy string
	// -n, --dry-run[=false]
	// do not delete anything, just print what would be done
	DryRun bool
	// --prune[=false]
	// automatically run the 'prune' command if snapshots have been removed
	Prune bool
	// --max-unused="5%"
	// tolerate  given  limit of unused data (absolute value in bytes with
	// suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited')
	MaxUnused string
	// --max-repack-size=""
	// maximum size to repack (allowed suffixes: k/K, m/M, g/G, t/T)
	MaxRepackSize string
	// --repack-cacheable-only[=false]
	// only repack packs which are cacheable
	RepackCacheableOnly bool
	// -h, --help[=false]
	// help for forget
	Help bool

	GlobalFlags
}

func concatForgetFlags() {}
