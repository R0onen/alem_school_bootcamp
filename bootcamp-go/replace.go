package bootcamp

func Replace(s string, old string, new string) string {
	if len(s) == 0 {
		return ""
	}

	ln := len(s)

	for i := 0; i <= ln-len(old); i++ {
		if s[i:i+len(old)] == old {
			s = s[:i] + new + s[i+len(old):]
			i += len(new) - 1
			ln = len(s)
		}
	}
	return s
}
