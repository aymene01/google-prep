package dsa

func TwoSum(values []int, target int) bool {
	seen := map[int]bool{}

	for _, num := range values {
		diff := target - num

		if seen[diff] {
			return true
		}

		seen[num] = true
	}
	return false
}

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := map[rune]int{}

	for _, ch := range s1 {
		count[ch]++
	}


	return true
}