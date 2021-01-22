package main

func closeStrings(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := map[rune]int{}
	se := map[int]int{}
	tm := map[rune]int{}
	te := map[int]int{}
	for _, c := range s {
		sm[c]++
	}
	for _, n := range sm {
		se[n]++
	}
	for _, c := range t {
		tm[c]++
	}
	for _, n := range tm {
		te[n]++
	}
	for k, v := range se {
		if te[k] != v {
			return false
		}
	}
	for k, v := range te {
		if se[k] != v {
			return false
		}
	}
	for k, _ := range sm {
		if _, ok := tm[k]; !ok {
			return false
		}
	}
	for k, _ := range tm {
		if _, ok := sm[k]; !ok {
			return false
		}
	}
	return true
}
