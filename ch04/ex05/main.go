package main

func main() {}

func removeDuplicate(ss []string) []string {
	dup := ss[0]
	for i, s := range ss[1:] {
		if s == dup {
			ss[i] = ""
		} else {
			dup = s
		}
	}
	return nonempty(ss)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func equals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
