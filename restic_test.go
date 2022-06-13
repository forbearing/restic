package restic

import (
	"context"
	"os"
	"testing"
)

// TODO:
// 测试 restic 命令, 这些命令需要一个 snapshot ID
//     ls, dump, restore, tag

func TestRestic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := "/tmp/restic_repo"
	backupSource1 := "./testdata/backup_source_1"
	backupSource2 := "./testdata/backup_source_2"
	os.RemoveAll(repo)
	os.Mkdir(repo, 0755)

	// global flags
	gopt := make(map[string]string)
	gopt["name"] = "restic"
	gopt["type"] = "backup"
	globalF := &GlobalFlags{
		Cacert:       []string{"./testdata/certs/example.com.crt", "./testdata/certs/nginx.example.com.crt"},
		CacheDir:     "/tmp/cache",
		CleanupCache: true,
		InsecureTls:  true,
		Json:         false,
		Option:       gopt,
		Repo:         repo,
	}
	// init command
	initC := &Init{}
	// backup command
	backupC := &Backup{
		Tag:  []string{"mytag", "test"},
		Host: "myhost",
	}
	backupC.SetArgs(backupSource1, backupSource2)
	// snapshots command
	snapshotsC := &Snapshots{}
	// list command
	listC := &List{}
	listC.SetArgs("snapshots")
	// cache command
	cacheC := &Cache{}
	// check command
	checkC := &Check{}
	// find command
	findC := &Find{}
	findC.SetArgs("*hosts")
	// forget command
	forgetC := &Forget{
		Tag:      []string{"mytag", "test"},
		Host:     []string{"myhost"},
		KeepLast: 3,
	}
	// generate command
	generateC := &Generate{
		BashCompletion: "/tmp/restic.bash",
		ZshCompletion:  "/tmp/restic.zsh",
		FishCompletion: "/tmp/restic.fish",
	}
	// key command
	keyC := &Key{}
	keyC.SetArgs("list")
	// migrate command
	migrateC := &Migrate{}
	_ = migrateC

	restic, err := New(ctx, globalF)
	if err != nil {
		t.Error(err)
	}

	t.Log(restic.String())
	t.Log(restic.Command(initC).String())
	t.Log(restic.Command(backupC).String())
	t.Log(restic.Command(snapshotsC).String())
	t.Log(restic.Command(listC).String())
	t.Log(restic.Command(cacheC).String())
	t.Log(restic.Command(checkC).String())
	t.Log(restic.Command(findC).String())
	t.Log(restic.Command(forgetC).String())
	t.Log(restic.Command(generateC).String())
	t.Log(restic.Command(keyC).String())
	t.Log(restic.Command(&Migrate{}).String())
	t.Log(restic.Command(&Prune{}).String())
	t.Log(restic.Command(&RebuildIndex{}).String())
	t.Log(restic.Command(&Recover{}).String())
	t.Log(restic.Command(&Stats{
		Tag:  []string{"mytag", "test"},
		Host: []string{"myhost"},
	}).String())
	t.Log(restic.Command(&Tag{}).String())
	t.Log(restic.Command(&Unlock{}).String())
	t.Log(restic.Command(&Version{}).String())

	restic.Run()
	restic.Command(initC).Run()
	restic.Command(backupC).Run()
	restic.Command(backupC).Run()
	restic.Command(backupC).Run()
	restic.Command(backupC).Run()
	restic.Command(snapshotsC).Run()
	restic.Command(listC).Run()
	restic.Command(cacheC).Run()
	restic.Command(checkC).Run()
	restic.Command(findC).Run()
	restic.Command(forgetC).Run()
	restic.Command(generateC).Run()
	restic.Command(keyC).Run()
	restic.Command(&Migrate{}).Run()
	restic.Command(&Prune{}).Run()
	restic.Command(&RebuildIndex{}).Run()
	restic.Command(&Recover{}).Run()
	restic.Command(&Stats{
		Tag:  []string{"mytag", "test"},
		Host: []string{"myhost"},
	}).Run()
	restic.Command(&Tag{}).Run()
	restic.Command(&Unlock{}).Run()
	restic.Command(&Version{}).Run()
}
