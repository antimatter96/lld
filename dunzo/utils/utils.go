package utils

func CopyMap(mp map[string]int) map[string]int {
	cp := make(map[string]int, len(mp))

	for k, v := range mp {
		cp[k] = v
	}

	return cp
}
