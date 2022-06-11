package restic

import (
	"testing"
)

func TestConcatGlobalFlag(t *testing.T) {
	option := make(map[string]string)
	option["name"] = "restic"
	option["type"] = "backup"
	f := &GlobalFlags{
		Cacert:       []string{"./testdata/certs/example.com.crt", "./testdata/certs/nginx.example.com.crt"},
		CacheDir:     "/tmp/cache",
		CleanupCache: true,
		InsecureTls:  true,
		Option:       option,
		Json:         false,
		//Repo:         "./testdata/repo",
	}

	s := concatAll(f)
	t.Log(s)
}
