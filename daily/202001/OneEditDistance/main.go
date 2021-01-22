package main

func isOneEditDistance(s string, t string) bool {
	switch len(s) - len(t) {
	case 0:
		//replace only
		distance := 0
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				distance++
			}
		}
		if distance == 1 {
			return true
		}
		return false
	case 1:
		for i := 0; i < len(s); i++ {
			ns := s[:i] + s[i+1:]
			if ns == t {
				return true
			}
		}
		return false
	case -1:
		for i := 0; i < len(t); i++ {
			nt := t[:i] + t[i+1:]
			if nt == s {
				return true
			}
		}
		return false
	default:
		return false
	}
}
