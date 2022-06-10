package restic

// - MigrateFlags includes all flags of "restic migrate" and inheris GlobalFlags.
// - The "migrate" command applies migrations to a repository. When no migration
//   name is explicitly given, a list of migrations that can be applies is printed.
type MigrateFlags struct {
	// -f, --force[=false]
	// apply a migration a second time
	Force bool
	// -h, --help[=false]
	// help for migrate
	Help bool

	GlobalFlags
}

func concatMigrateFlags() {}
