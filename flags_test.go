package restic

import "testing"

func TestConcatGlobalFlag(t *testing.T) {
	f := new(GlobalFlags)
	f.Cacert = []string{"/etc/cacert.pem"}

	s := ConcatFlags(f)
	t.Log(s)
}
