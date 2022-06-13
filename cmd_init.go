package restic

import "strings"

// - Init includes all flags of "restic init" and inheris GlobalFlags.
// - The "init" command initializes a new repository.
type Init struct {
	// --copy-chunker-params[=false]
	// copy chunker parameters from the secondary repository (useful with the copy command)
	CopyChunkerParams bool `json:"--copy-chunker-params"`
	// -h, --help[=false]
	// help for init
	Help bool `json:"--help"`
	// --key-hint2=""
	// key ID of key to try decrypting the secondary repository first
	// (default: $RESTIC_KEY_HINT2)
	KeyHint2 string `json:"--key-hint2"`
	// --password-command2=""
	// shell command to obtain the secondary repository password from
	// (default: $RESTIC_PASSWORD_COMMAND2)
	PasswordCommand2 string `json:"--password-command2"`
	// --password-file2=""
	// file to read the secondary repository password from
	// (default: $RESTIC_PASSWORD_FILE2)
	PasswordFile2 string `json:"--password-file2"`
	// --repo2=""
	// secondary repository to copy chunker parameters from
	// (default: $RESTIC_REPOSITORY2)
	Repo2 string `json:"--repo2"`
	// --repository-file2=""
	// file from which to read the secondary repository location to copy
	// chunker parameters from (default: $RESTIC_REPOSITORY_FILE2)
	RepositoryFile2 string `json:"--repository-file2"`

	args strings.Builder

	GlobalFlags
}

func (i *Init) Flags() string {
	return concat(i)
}

func (i *Init) Name() string {
	return "init"
}

func (i *Init) Args() string {
	return strings.TrimSpace(i.args.String())
}

func (i *Init) SetArgs(args ...string) string {
	for _, s := range args {
		i.args.WriteString(s + " ")
	}
	return i.Args()
}
