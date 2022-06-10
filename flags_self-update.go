package restic

// - SelfUpdateFlags includes all flags of "restic self-update"
//   and inheris GlobalFlags
// - The "self-update" command downloads the latest stable release of restic
//   from Github and replaces the currently running binary. After download,
//   the authenticity of the binary is verified using the GPG signature on
//   the release files.
type SelfUpdateFlags struct {
	// -h, --help[=false]
	// help for self-update
	Help bool
	// --output=""
	// Save the downloaded file as filename (default: running binary itself)
	Output string

	GlobalFlags
}

func concatSelfUpdateFlags() {}
