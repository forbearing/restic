package restic

func ConcatFlags(fl ...Flag) string {
	var s string
	for _, f := range fl {
		s = s + f.Concat()
	}
	return s
}
