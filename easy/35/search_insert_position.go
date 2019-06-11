package searchInsertPosition

func searchInsert(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s < e {
		m := (s + e) / 2
		if nums[m] == target {
			return m
		} else if target < nums[m] {
			e = m
		} else {
			s = m + 1
		}
	}
	if target > nums[e] {
		return e + 1
	}
	return s
}
