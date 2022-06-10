package restic

// - RecoverFlags includes all flags of "restic recover" and inheris GlobalFlags
// - The "recover" command builds a new repository from all directories it can find
//   in the raw data of the repository which are not referenced in an existing
//   snapshot. It can used if, for example, a snapshot has been removed by accident
//   with "forget"
type RecoverFlags struct {
	// -h, --help[=false]      help for recover
	Help bool

	GlobalFlags
}

func concatRecoverFlags() {}
