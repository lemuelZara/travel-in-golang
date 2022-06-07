package basics

func Sum(nums []int) (result int) {
	for _, num := range nums {
		result += num
	}

	return result
}

func SumAll(numsToSum ...[]int) (result []int) {
	result = make([]int, 0, len(numsToSum))

	for _, nums := range numsToSum {
		result = append(result, Sum(nums))
	}

	return result
}
