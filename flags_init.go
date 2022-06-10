package restic

// - InitFlags includes all flags of "restic init" and inheris GlobalFlags.
// - The "init" command initializes a new repository.
type InitFlags struct {
	// --copy-chunker-params[=false]
	// copy chunker parameters from the secondary repository (useful with the copy command)
	CopyChunkerParams bool
	// -h, --help[=false]
	// help for init
	Help bool
	// --key-hint2=""
	// key ID of key to try decrypting the secondary repository first
	// (default: $RESTIC_KEY_HINT2)
	KeyHint2 string
	// --password-command2=""
	// shell command to obtain the secondary repository password from
	// (default: $RESTIC_PASSWORD_COMMAND2)
	PasswordCommand2 string
	// --password-file2=""
	// file to read the secondary repository password from
	// (default: $RESTIC_PASSWORD_FILE2)
	PasswordFile2 string
	// --repo2=""
	// secondary repository to copy chunker parameters from
	// (default: $RESTIC_REPOSITORY2)
	Repo2 string
	// --repository-file2=""
	// file from which to read the secondary repository location to copy
	// chunker parameters from (default: $RESTIC_REPOSITORY_FILE2)
	RepositoryFile2 string

	GlobalFlags
}

func concatInitFlags() {}
