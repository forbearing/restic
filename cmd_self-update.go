package restic

import "strings"

// - SelfUpdate includes all flags of "restic self-update"
//   and inheris GlobalFlags
// - The "self-update" command downloads the latest stable release of restic
//   from Github and replaces the currently running binary. After download,
//   the authenticity of the binary is verified using the GPG signature on
//   the release files.
type SelfUpdate struct {
	// -h, --help[=false]
	// help for self-update
	Help bool `json:"--help"`
	// --output=""
	// Save the downloaded file as filename (default: running binary itself)
	Output string `json:"--output"`

	args strings.Builder

	GlobalFlags
}

func (s *SelfUpdate) Flags() string {
	return concat(s)
}

func (s *SelfUpdate) Name() string {
	return "self-update"
}

func (s *SelfUpdate) Args() string {
	return strings.TrimSpace(s.args.String())
}

func (s *SelfUpdate) SetArgs(args ...string) string {
	for _, str := range args {
		s.args.WriteString(str + " ")
	}
	return s.Args()
}
