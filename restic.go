package restic

import (
	"context"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type Flag interface {
	Concat() string
}

type Restic struct {
	cmdName   string
	cmdArgs   string
	CmdString string
	cmd       *exec.Cmd

	finished bool
	ctx      context.Context
	waitDone chan struct{}
	l        sync.Mutex
}

// New
func New(ctx context.Context, fl ...Flag) (*Restic, error) {
	r := new(Restic)

	path, err := exec.LookPath("restic")
	if err != nil {
		return nil, err
	}
	r.cmdName = filepath.Base(path)

	// concat all restic command and sub-commands flags
	for _, f := range fl {
		r.cmdArgs = r.cmdArgs + f.Concat()
	}
	r.CmdString = r.cmdName + r.cmdArgs

	return r, nil
}

// Run
func (r *Restic) Run() error {
	var (
		done = make(chan struct{}, 1)
		// done is a channel that wait goroutine to output stdout and stderr.
		stdout io.ReadCloser
		stderr io.ReadCloser
		err    error
	)
	r.cmd = exec.Command(r.cmdName, strings.Fields(r.cmdArgs)...)
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
