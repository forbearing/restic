## Intruduction

A wrapper for the backup tools restc

## Installation

` go get github.com/forbearing/restic@v1.1.0`

## How to use this lib



```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/forbearing/restic"
)

var (
	ctx, cancel   = context.WithTimeout(context.Background(), time.Minute*10)
	backupSource1 = "./testdata/backup_source_1"
	backupSource2 = "./testdata/backup_source_2"
)

func main() {
	defer cancel()

	// new a restic instance.
	r, err := restic.New(ctx, &restic.GlobalFlags{
		NoCache:       true,
		LimitDownload: 20480,
		LimitUpload:   1024,
	})
	if err != nil {
		log.Fatal(err)
	}

	// restic backup
	r.Command(restic.Backup{
		Host: "myhost",
		Tag:  []string{"mytag", "test"},
	}.SetArgs(backupSource1, backupSource2))
	fmt.Println(r)
	//r.Run() // execute `restic backup command`

	// restic backup
	r.Command(&restic.Snapshots{
		Host: []string{"myhost"},
		Tag:  []string{"mytag"},
	})
	fmt.Println(r)
	//r.Run() // execute `restic snapshots command`

	// restic snapshots
	r.Command(&restic.Forget{
		KeepLast: 1,
	})
	fmt.Println(r)
	//r.Run() // execute `restic forget command`

	// Output:

	//restic --limit-download=20480 --limit-upload=1024 --no-cache backup --host=myhost --tag=mytag,test ./testdata/backup_source_1 ./testdata/backup_source_2
	//restic --limit-download=20480 --limit-upload=1024 --no-cache snapshots --host=myhost --tag=mytag
	//restic --limit-download=20480 --limit-upload=1024 --no-cache forget --keep-last=1
}
```



