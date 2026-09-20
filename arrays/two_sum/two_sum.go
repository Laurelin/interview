package twosum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if index, ok := seen[complement]; ok {
			return []int{index, i}
		}
		seen[nums[i]] = i
	}
    return nil
}