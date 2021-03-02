package utils

func ShowDiff(strA, strB []string) (strAOnly []string, strBOnly []string, both []string) {
	mapA := make(map[string]struct{})
	for _, s := range strA {
		mapA[s] = struct{}{}
	}
	mapB := make(map[string]struct{})
	for _, s := range strB {
		if _, ok := mapA[s]; !ok {
			strBOnly = append(strBOnly, s)
		} else {
			both = append(both, s)
		}
		mapB[s] = struct{}{}
	}
	for _, s := range strA {
		if _, ok := mapB[s]; !ok {
			strAOnly = append(strAOnly, s)
		}
	}
	return
}
