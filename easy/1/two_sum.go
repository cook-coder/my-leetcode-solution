package twoSum

func twoSumBadWay(nums []int, target int) []int {
	indices := make([]int, 2, 2)
	for i, num := range nums {
		if num > target {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				indices[0] = i
				indices[1] = j
				return indices
			}
		}
	}
	return indices
}

func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}
	length := len(nums)
	for i := 0; i < length; i++ {
		numsMap[nums[i]] = i
	}
	for i := 0; i < length; i++ {
		complement := target - nums[i]
		if _, ok := numsMap[complement]; ok && numsMap[complement] != i {
			return []int{i, numsMap[complement]}
		}
	}
	return []int{}
}
