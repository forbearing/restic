package restic

import (
	"context"
	"os"
	"testing"
)

func TestRestic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := "/tmp/restic_repo"
	backupSource := "./testdata/backup_source"
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

	// init flags
	initF := &InitFlags{}

	// backup flags
	backupF := &BackupFlags{
		Tag:  []string{"mytag", "test"},
		Host: "myhost",
	}
	backupF.SetArgs(backupSource)

	// snapshots flags
	snapshotsF := &SnapshotsFlags{}

	r, err := New(ctx, globalF, initF)
	if err != nil {
		t.Error(err)
	}
	t.Log(r.String())
	r.Run()

	r1, err := New(ctx, globalF, backupF)
	if err != nil {
		t.Error(err)
	}
	t.Log(r1.String())
	r1.Run()
	r1.Run()
	r1.Run()
	r1.Run()
	r1.Run()

	r2, err := New(ctx, globalF, snapshotsF)
	if err != nil {
		t.Error(err)
	}
	t.Log(r2.String())
	r2.Run()
}
