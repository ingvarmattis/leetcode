package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)-1)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}

		m[num] = i
	}

	return []int{0, 0}
}
