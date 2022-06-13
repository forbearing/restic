package restic

import "strings"

// - Key includes all flags of "restic key" and inheris GlobalFlags
// - The "key" command manages keys(passwords) for accessing the repository
type Key struct {
	// -h, --help[=false]
	// help for key
	Help bool `json:"--help"`
	// --host=""
	// the hostname for new keys
	Host string `json:"--host"`
	// --new-password-file=""
	// file from which to read the new password
	NewPasswordFile string `json:"--new-password-file"`
	// --user=""
	// the username for new keys
	User string `json:"--user"`

	args strings.Builder

	GlobalFlags
}

func (k *Key) Flags() string {
	return concat(k)
}

func (k *Key) Name() string {
	return "key"
}

func (k *Key) Args() string {
	return strings.TrimSpace(k.args.String())
}

func (k *Key) SetArgs(args ...string) string {
	for _, s := range args {
		k.args.WriteString(s + " ")
	}
	return k.Args()
}
