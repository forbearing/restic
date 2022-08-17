package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/forbearing/restic"
)

func main() {
	repo := "/tmp/restic-repo"
	err := os.RemoveAll(repo)
	if err != nil {
		log.Println("remove /tmp/restic failed: ", err)
	}
	r := restic.NewIgnoreNotFound(context.Background(), &restic.GlobalFlags{Repo: repo, Quiet: true})
	fmt.Println(r)
	r.SetEnv("RESTIC_PASSWORD", "mypass")
	r.Command(restic.Init{}).Run()
}
