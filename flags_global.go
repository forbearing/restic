package restic

import (
	"fmt"
	"reflect"
)

// - GlobalFlags includes all restic global flags.
// - Restic is a backup system which allows saving multiple revisions of files
//  and directories in an encrypted repository stored on different backends.
type GlobalFlags struct {
	// --cacert=[]
	// file to load root certificates from (default: use system certificates)
	Cacert []string
	// --cache-dir=""
	// set the cache directory. (default: use system default cache directory)
	CacheDir string
	// --cleanup-cache[=false]
	// auto remove old cache directories
	CleanupCache bool
	// -h, --help[=false]
	// help for restic
	Help bool
	// --insecure-tls[=false]
	// skip TLS certificate verification when connecting to the repo (insecure)
	InsecureTls bool
	// --json[=false]
	// set output mode to JSON for commands that support it
	Json bool
	// --key-hint=""
	// key ID of key to try decrypting first (default: $RESTIC_KEY_HINT)
	KeyHint string
	// --limit-download=0
	// limits downloads to a maximum rate in KiB/s. (default: unlimited)
	LimitDownload int64
	// --limit-upload=0
	// limits uploads to a maximum rate in KiB/s. (default: unlimited)
	LimitUpload int64
	// --no-cache[=false]
	// do not use a local cache
	NoCache bool
	// --no-lock[=false]
	// do not lock the repository, this allows some operations on read-only repositories
	NoLock bool
	// -o, --option=[]
	// set extended option (key=value, can be specified multiple times)
	Option map[string]string
	// --password-command=""
	// shell command to obtain the repository password from (default: $RESTIC_PASSWORD_COMMAND)
	PasswordCommand string
	// -p, --password-file=""
	// file to read the repository password from (default: $RESTIC_PASSWORD_FILE)
	PasswordFile string
	// -q, --quiet[=false]
	// do not output comprehensive progress report
	Quiet bool
	// -r, --repo=""
	// repository to backup to or restore from (default: $RESTIC_REPOSITORY)
	Repo string
	// --repository-file=""
	// file to read the repository location from (default: $RESTIC_REPOSITORY_FILE)
	RepositoryFile string
	// --tls-client-cert=""
	// path to a file containing PEM encoded TLS client certificate and private key
	TlsClientCert string
	// -v, --verbose[=0]
	// be verbose (specify multiple times or a level using --verbose=n, max level/times is 3)
	Verbose int
}

// ref:
//     https://stackoverflow.com/questions/42294015/how-to-use-go-reflection-pkg-to-get-the-type-of-a-slice-struct-field
// Concat implements interface Flag
func (f *GlobalFlags) Concat() string {
	var s string
	v := reflect.ValueOf(f).Elem()

	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		field := v.Field(i)
		name := v.Type().Field(i).Name
		kind := v.Field(i).Kind()
		typ := field.Type().String()
		fmt.Printf("Name: %s  Kind: %s  Type: %s  Value: %s\n", name, kind, typ, val)

		//switch typ {
		//case "string":
		//case "[]string":
		//case "int":
		//case "int64":
		//case "bool":
		//case "map[string]string":
		//}
	}

	return s
}
