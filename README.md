## Intruduction

A wrapper for the backup tools restc

## Installation

` go get github.com/forbearing/restic@v1.1.1`

## How to use this lib



```go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/forbearing/restic"
	"github.com/sirupsen/logrus"
)

var (
	ctx, cancel   = context.WithTimeout(context.Background(), time.Minute*10)
	backupSource1 = "../testdata/backup_source_1"
	backupSource2 = "../testdata/backup_source_2"
	repo          = "/tmp/restic_repo"
)

func main() {
	defer cancel()

	// new a restic instance.
	r, err := restic.New(ctx, &restic.GlobalFlags{
		NoCache:       true,
		LimitDownload: 20480,
		LimitUpload:   1024,
		Repo:          repo,
	})
	if err != nil {
		log.Fatal(err)
	}
	writer := logrus.New().WriterLevel(logrus.ErrorLevel)
	// discard all standard output and output error message to logrus's writer.
	r.SetOutput(io.Discard, writer)

	// execute `restic init` command
	os.RemoveAll(repo)
	r.Command(restic.Init{}).Run()
	fmt.Println(r)

	//execute `restic backup` command
	r.Command(restic.Backup{
		Host: "myhost",
		Tag:  []string{"mytag", "test"},
	}.SetArgs(backupSource1, backupSource2)).Run()
	fmt.Println(r)
	r.Run() // backup data again.
	r.Run() // backup data again and again.

	// execute `restic snapshots` command
	r.Command(&restic.Snapshots{
		Host: []string{"myhost"},
		Tag:  []string{"mytag"},
	}).Run()
	fmt.Println(r)

	// Output:

	//restic --limit-download=20480 --limit-upload=1024 --no-cache --repo=/tmp/restic_repo init
	//restic --limit-download=20480 --limit-upload=1024 --no-cache --repo=/tmp/restic_repo backup --host=myhost --tag=mytag,test ../testdata/backup_source_1 ../testdata/backup_source_2
	//restic --limit-download=20480 --limit-upload=1024 --no-cache --repo=/tmp/restic_repo snapshots --host=myhost --tag=mytag
}
```



