package dsa


func ContainDuplicate(array []int) bool {
	seen := map[int]bool{}

	for _, num := range array {
		if seen[num] {
			return true
		}

		seen[num] = true
	}
	return false
}