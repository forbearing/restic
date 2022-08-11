package restic

import "strings"

// - GlobalFlags includes all restic global flags.
// - Restic is a backup system which allows saving multiple revisions of files
//   and directories in an encrypted repository stored on different backends.
type GlobalFlags struct {
	// --cacert=[]
	// file to load root certificates from (default: use system certificates)
	Cacert []string `json:"--cacert"`
	// --cache-dir=""
	// set the cache directory. (default: use system default cache directory)
	CacheDir string `json:"--cache-dir"`
	// --cleanup-cache[=false]
	// auto remove old cache directories
	CleanupCache bool `json:"--cleanup-cache"`
	// -h, --help[=false]
	// help for restic
	Help bool `jons:"--help"`
	// --insecure-tls[=false]
	// skip TLS certificate verification when connecting to the repo (insecure)
	InsecureTls bool `json:"--insecure-tls"`
	// --json[=false]
	// set output mode to JSON for commands that support it
	Json bool `json:"--json"`
	// --key-hint=""
	// key ID of key to try decrypting first (default: $RESTIC_KEY_HINT)
	KeyHint string `json:"--key-hint"`
	// --limit-download=0
	// limits downloads to a maximum rate in KiB/s. (default: unlimited)
	LimitDownload int64 `json:"--limit-download"`
	// --limit-upload=0
	// limits uploads to a maximum rate in KiB/s. (default: unlimited)
	LimitUpload int64 `json:"--limit-upload"`
	// --no-cache[=false]
	// do not use a local cache
	NoCache bool `json:"--no-cache"`
	// --no-lock[=false]
	// do not lock the repository, this allows some operations on read-only repositories
	NoLock bool `json:"--no-lock"`
	// -o, --option=[]
	// set extended option (key=value, can be specified multiple times)
	Option map[string]string `json:"--option"`
	// --password-command=""
	// shell command to obtain the repository password from (default: $RESTIC_PASSWORD_COMMAND)
	PasswordCommand string `json:"--password-command"`
	// -p, --password-file=""
	// file to read the repository password from (default: $RESTIC_PASSWORD_FILE)
	PasswordFile string `json:"--passwod-file"`
	// -q, --quiet[=false]
	// do not output comprehensive progress report
	Quiet bool `json:"--quiet"`
	// -r, --repo=""
	// repository to backup to or restore from (default: $RESTIC_REPOSITORY)
	Repo string `json:"--repo"`
	// --repository-file=""
	// file to read the repository location from (default: $RESTIC_REPOSITORY_FILE)
	RepositoryFile string `json:"--repository-file"`
	// --tls-client-cert=""
	// path to a file containing PEM encoded TLS client certificate and private key
	TlsClientCert string `json:"--tls-client-cert"`
	// -v, --verbose[=0]
	// be verbose (specify multiple times or a level using --verbose=n, max level/times is 3)
	Verbose int `json:"--verbose"`

	args strings.Builder
}

// ref:
//     https://stackoverflow.com/questions/42294015/how-to-use-go-reflection-pkg-to-get-the-type-of-a-slice-struct-field
// Concat implements interface Flag
func (g *GlobalFlags) Flags() string {
	return concat(g)
}

func (g *GlobalFlags) Name() string {
	return ""
}

func (g *GlobalFlags) Args() string {
	return strings.TrimSpace(g.args.String())
}

//func (g *GlobalFlags) SetArgs(args ...string) string {
//    for _, s := range args {
//        g.args.WriteString(s + " ")
//    }
//    return g.Args()
//}
