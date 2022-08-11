package restic

import "strings"

// - Copy includes all flags of "restic copy" and inheris GlobalFlags.
// - The "Copy" command copies one or more snapshots from one repository to another.
// - NOTE:  This  process  will have to both download (read) and upload (write)
//   the entire snapshot(s) due to the different encryption keys used in the source
//   and destination repositories. This /may incur higher bandwidth usage and
//   costs/ than expected during normal backup runs.
// - NOTE: The copying process does not re-chunk files, which may break deduplication
//   between the files copied and files already  stored  in  the  destination repository.
//   This means that copied files, which existed in both the source and destination
//   repository, /may occupy up to twice their space/ in the destination repository.
//   This can be mitigated by the "--copy-chunker-params" option when initializing
//   a new destination repository using the "init" command.
type Copy struct {
	// -h, --help[=false]
	// help for copy
	Help bool `json:"--help"`
	// -H, --host=[]
	// only consider snapshots for this host, when no snapshot ID is given
	// (can be specified multiple times)
	Host []string `json:"--host"`

	// --key-hint2=""
	// key ID of key to try decrypting the destination repository first
	// (default: $RESTIC_KEY_HINT2)
	KeyHint2 string `json:"--key-hint2"`
	// --password-command2=""
	// shell command to obtain the destination repository password from
	// (default: $RESTIC_PASSWORD_COMMAND2)
	PasswordCommand2 string `json:"--password-command2"`
	// --password-file2=""
	// file to read the destination repository password from
	// (default: $RESTIC_PASSWORD_FILE2)
	PasswordFile2 string `json:"--password-file2"`
	// --path=[]
	// only consider snapshots which include this (absolute) path,
	// when no snapshot ID is given
	Path []string `json:"--path"`
	// --repo2=""
	// destination repository to copy snapshots to (default: $RESTIC_REPOSITORY2)
	Repo2 string `json:"--repo2"`
	// --repository-file2=""
	// file from which to read the destination repository location to copy
	// snapshots to (default: $RESTIC_REPOSITORY_FILE2)
	RepositoryFile2 string `json:"--repository-file2"`
	// --tag=[]
	// only consider snapshots which include this taglist,
	// when no snapshot ID is given
	Tag string `json:"--tag"`

	args strings.Builder

	GlobalFlags
}

func (c Copy) Name() string  { return "copy" }
func (c Copy) Flags() string { return concatFlags(c) }
func (c Copy) Args() string  { return strings.TrimSpace(c.args.String()) }

func (c Copy) SetArgs(args ...string) *Copy {
	for _, s := range args {
		c.args.WriteString(s + " ")
	}
	return &c
}
