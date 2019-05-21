package twoSum

import "testing"

func TestTwoSum(t *testing.T) {
	//nums := []int{2, 7, 11, 15}
	//target := 9
	nums := []int{3, 2, 4}
	target := 6

	r := twoSum(nums, target)
	t.Log(r)
}
