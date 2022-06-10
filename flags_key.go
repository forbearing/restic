package restic

// - KeyFlags includes all flags of "restic key" and inheris GlobalFlags
// - The "key" command manages keys(passwords) for accessing the repository
type KeyFlags struct {
	// -h, --help[=false]
	// help for key
	Help bool
	// --host=""
	// the hostname for new keys
	Host string
	// --new-password-file=""
	// file from which to read the new password
	NewPasswordFile string
	// --user=""
	// the username for new keys
	User string

	GlobalFlags
}

func concatKeyFlags() {}
