package restic

import (
	"context"
	"testing"
)

func TestRestic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	option := make(map[string]string)
	option["name"] = "restic"
	option["type"] = "backup"
	f := &GlobalFlags{
		Cacert:       []string{"./testdata/certs/example.com.crt", "./testdata/certs/nginx.example.com.crt"},
		CacheDir:     "/tmp/cache",
		CleanupCache: true,
		InsecureTls:  true,
		Json:         false,
		Option:       option,
	}

	r, err := New(ctx, f)
	if err != nil {
		t.Error(err)
	}
	t.Log(r.CmdString)
}
