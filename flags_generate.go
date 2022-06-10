package restic

// - GenerateFlags includes all flags of "restic generate" and inheris GlobalFlags
// - The "generate" command writes automatically generated files (like the man pages
//   and the auto-completion files for bash, fish and zsh).
type GenerateFlags struct {
	// --bash-completion=""
	// write bash completion file
	BashCompletion string
	// --fish-completion=""
	// write fish completion file
	FishCompletion string
	// -h, --help[=false]
	// help for generate
	Help bool
	// --man=""
	// write man pages to directory
	Man string
	// --zsh-completion=""
	// write zsh completion file
	ZshCompletion string

	GlobalFlags
}

func concatGenerateFlags() {}
