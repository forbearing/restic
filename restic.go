package restic

import (
	"context"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// Name return command name.
// Flags return the command all concatenated flags.
// Args return the command all arguments.
type Command interface {
	Name() string
	Flags() string
	Args() string
}

type Restic struct {
	resticName  string // like "restic", "restic_darwin_amd64"
	globalFlags string // restic global flags
	cmdName     string // restic sub-command name
	cmdFlags    string // restic sub-command flags
	cmdArgs     string // restic sub-command arguments

	cmd *exec.Cmd

	finished bool
	ctx      context.Context
	waitDone chan struct{}
	l        sync.Mutex
}

// New returns a restic instance.
func New(ctx context.Context, g *GlobalFlags) (*Restic, error) {
	r := new(Restic)

	path, err := exec.LookPath("restic")
	if err != nil {
		return nil, err
	}
	r.resticName = filepath.Base(path)

	r.globalFlags = g.Flags()

	return r, nil
}

func (r *Restic) Command(c Command) *Restic {
	r.cmdName = c.Name()
	r.cmdFlags = c.Flags()
	r.cmdArgs = c.Args()
	return r
}

// String returns restic commmand line
// such like "restic --limit-upload=0 -v=0 snapshots --tag=mytag --host=myhost"
func (r *Restic) String() string {
	builder := new(strings.Builder)

	builder.WriteString(r.resticName + " ")
	builder.WriteString(r.globalFlags + " ")
	builder.WriteString(r.cmdName + " ")
	// If r.cmdFlags is empty string, builder.WriteString will add one more
	// space to restic command line.
	// It necessary to ignore it when r.cmdFlags is empty.
	if len(r.cmdFlags) != 0 {
		builder.WriteString(r.cmdFlags + " ")
	}
	// cmdArgs is the same as r.cmdFlags.
	if len(r.cmdArgs) != 0 {
		builder.WriteString(r.cmdArgs)
	}

	return builder.String()
}

// Run start execute restic command line
// restic command line string returned by Restic.String() method.
func (r *Restic) Run() error {
	var (
		// done is a channel that wait goroutine to output stdout and stderr.
		done   = make(chan struct{}, 1)
		stdout io.ReadCloser
		stderr io.ReadCloser
		err    error
	)

	cmdString := strings.Fields(r.String())
	r.cmd = exec.Command(cmdString[0], cmdString[1:]...)
	if stdout, err = r.cmd.StdoutPipe(); err != nil {
		return err
	}
	if stderr, err = r.cmd.StderrPipe(); err != nil {
		return err
	}

	if err = r.cmd.Start(); err != nil {
		return err
	}
	print(stdout, stderr, done)
	<-done
	if err = r.cmd.Wait(); err != nil {
		return err
	}

	return nil
}
