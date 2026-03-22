package utils

func Sum(nums ...int) int {
	if len(nums) <= 0 || len(nums) > 50 {
		return 0
	}
	total := 0
	for _, v := range nums {
		total += v
	}

	return total

}
