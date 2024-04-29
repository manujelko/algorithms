package algorithms

func LinearSearch(needle int, haystack []int) bool {
	for _, e := range haystack {
		if needle == e {
			return true
		}
	}
	return false
}
