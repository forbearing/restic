package restic

import (
	"context"
	"io"
	"os"
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
	resticName  string    // like "restic", "restic_darwin_amd64"
	globalFlags string    // restic global flags
	cmdName     string    // restic sub-command name
	cmdFlags    string    // restic sub-command flags
	cmdArgs     string    // restic sub-command arguments
	cmdPrefix   string    // add a prefix string before restic output with every line
	cmdStdout   io.Writer // cmdStdout is an io.Writer where restic command line normal output writes to
	cmdStderr   io.Writer // cmdStderr is an io.Writer where restic command line error output writes to

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

// Command setup restic commmand name, command flags and command arguments.
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

// SetOutput setup the restic command line normal output and error output.
// stdout is an io.Writer where restic command line normal output writes to.
// stderr is an io.Writer where restic command line error output writes to.
// Either stdout or stderr is nil, the restic command line output still is
// os.Stdout and os.Stderr.
// If not call SetOutput, the default output is alas os.Stdout and os.Stderr.
func (r *Restic) SetOutput(stdout, stderr io.Writer) {
	r.cmdStdout = stdout
	r.cmdStderr = stderr
}

// Run start execute restic command line
// restic command line string returned by Restic.String() method.
func (r *Restic) Run() error {
	cmdString := strings.Fields(r.String())
	r.cmd = exec.Command(cmdString[0], cmdString[1:]...)

	// setup r.cmd's stdout and stderr
	if r.cmdStdout != nil && r.cmdStderr != nil {
		r.cmd.Stdout = r.cmdStdout
		r.cmd.Stderr = r.cmdStderr
	} else {
		r.cmd.Stdout = os.Stdout
		r.cmd.Stderr = os.Stderr
	}

	if err := r.cmd.Start(); err != nil {
		return err
	}
	if err := r.cmd.Wait(); err != nil {
		return err
	}

	return nil
}
