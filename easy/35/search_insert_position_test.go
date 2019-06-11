package searchInsertPosition

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{1, 3, 4, 5, 6, 7, 9}
	//nums = []int{1}
	index := searchInsert(nums, 5)
	t.Log(index)
}
