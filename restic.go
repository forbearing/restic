package restic

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
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

func New(ctx context.Context) (*Restic, error) {
	r := new(Restic)

	path, err := exec.LookPath("restic")
	if err != nil {
		return nil, err
	}
	r.cmdName = filepath.Base(path)

	r.CmdString = r.cmdName + r.cmdArgs

	return r, nil
}

// print is used to output stdout and stderr in real time.
func print(stdout, stderr io.ReadCloser, done chan struct{}) {
	stopCh := make(chan struct{}, 2)
	defer stdout.Close()
	defer stderr.Close()
	//errCh := make(chan error, 2)

	// A goroutine that outputs stdout in real time.
	go func() {
		defer func() {
			stopCh <- struct{}{}
		}()
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Printf("%s", scanner.Text())
		}
		err := scanner.Err()
		// if stdout already closed, stop the goroutine.
		if errors.Is(err, os.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Println("scanner output stdout error:", err)
			return
		}
	}()

	// A goroutine that outputs stderr in real time.
	go func() {
		defer func() {
			stopCh <- struct{}{}
		}()
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Printf("%s", scanner.Text())
		}
		err := scanner.Err()
		// if stderr already closed, stop the goroutine.
		if errors.Is(err, os.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Println("scanner output stderr error:", err)
			return
		}
	}()

	<-stopCh
	<-stopCh
	done <- struct{}{}
}
