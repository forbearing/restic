package restic

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// TODO:
// 测试 restic 命令, 这些命令需要一个 snapshot ID
//     ls, dump, restore, tag

func TestRestic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := "/tmp/restic_repo"
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
	restic, err := New(ctx, globalF)
	if err != nil {
		t.Error(err)
	}
	notice := color.New(color.Bold, color.FgGreen).PrintfFunc()

	notice("\n\n----- test restic default output.\n\n")
	testRestic(t, restic)

	notice("\n\n----- test restic output to custom logger.\n\n")
	l := logrus.New()
	logw := l.Writer()
	defer logw.Close()
	restic.SetOutput(logw, logw)
	testRestic(t, restic)

	notice("\n\n----- test restic output to os.Stdout and os.Stderr.\n\n")
	restic.SetOutput(os.Stdout, os.Stderr)
	testRestic(t, restic)

	stdoutFile, err := os.Create("/tmp/restic_stdout.log")
	if err != nil {
		t.Fatal("os.Create /tmp/restic_stdout.log error:", err)
	}
	defer stdoutFile.Close()
	stderrFile, err := os.Create("/tmp/restic_stderr.log")
	if err != nil {
		t.Fatal("os.Create restic_stderr.log error:", err)
	}
	defer stderrFile.Close()
	restic.SetOutput(stdoutFile, stderrFile)
	notice("\n\n----- test restic output to file.\n\n")
	testRestic(t, restic)

	notice("\n\n----- test run restic command with timeoutable context.\n\n")
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel2()
	restic2, err := New(ctx2, globalF)
	if err != nil {
		t.Error(err)
	}
	restic2.SetOutput(os.Stdout, os.Stderr)
	testRestic(t, restic2)
}

func testRestic(t *testing.T, restic *Restic) {
	backupSource1 := "./testdata/backup_source_1"
	backupSource2 := "./testdata/backup_source_2"

	// init command
	initC := &Init{}
	// backup command
	backupC := &Backup{
		Tag:  []string{"mytag", "test"},
		Host: "myhost",
	}
	backupC.SetArgs(backupSource1, backupSource2)
	// snapshots command
	snapshotsC := &Snapshots{Latest: 1}
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

	notice := color.New(color.Bold, color.FgYellow).PrintlnFunc()

	notice(restic.String())
	restic.Run()

	// restic init
	notice(restic.Command(initC).String())
	restic.Command(initC).Run()

	// restic backup
	notice(restic.Command(backupC).String())
	restic.Command(backupC).Run()
	restic.Command(backupC).Run()
	restic.Command(backupC).Run()
	restic.Command(backupC).Run()

	// rsetic snapshots
	notice(restic.Command(snapshotsC).String())
	restic.Command(snapshotsC).Run()

	// restic list
	notice(restic.Command(listC).String())
	restic.Command(listC).Run()

	// restic cache
	notice(restic.Command(cacheC).String())
	restic.Command(cacheC).Run()

	// restic check
	notice(restic.Command(checkC).String())
	restic.Command(checkC).Run()

	// restic find
	notice(restic.Command(findC).String())
	restic.Command(findC).Run()

	// restic forget
	notice(restic.Command(forgetC).String())
	restic.Command(forgetC).Run()

	// restic generate
	notice(restic.Command(generateC).String())
	restic.Command(generateC).Run()

	// restic key
	notice(restic.Command(keyC).String())
	restic.Command(keyC).Run()

	// restic migrate
	notice(restic.Command(&Migrate{}).String())
	restic.Command(&Migrate{}).Run()

	// restic prune
	notice(restic.Command(&Prune{}).String())
	restic.Command(&Prune{}).Run()

	// restic rebuild-index
	notice(restic.Command(&RebuildIndex{}).String())
	restic.Command(&RebuildIndex{}).Run()

	// restic recover
	notice(restic.Command(&Recover{}).String())
	restic.Command(&Recover{}).Run()

	// restic tag
	notice(restic.Command(&Tag{}).String())
	restic.Command(&Tag{}).Run()

	// restic stats
	notice(restic.Command(&Stats{
		Tag:  []string{"mytag", "test"},
		Host: []string{"myhost"},
	}).String())
	restic.Command(&Tag{}).Run()
	restic.Command(&Stats{
		Tag:  []string{"mytag", "test"},
		Host: []string{"myhost"},
	}).Run()

	// restic unlock
	notice(restic.Command(&Unlock{}).String())
	restic.Command(&Unlock{}).Run()

	// restic version
	notice(restic.Command(&Version{}).String())
	restic.Command(&Version{}).Run()
}
