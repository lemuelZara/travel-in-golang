package basics

func RepeatWordAtTimes(word string, times int) string {
	var result string

	for i := 0; i < times; i++ {
		result += word
	}

	return result
}

func FindEvenNumbersByRange(start, end int) []int {
	var n []int

	for i := start; i <= end; i++ {
		if i%2 == 0 {
			n = append(n, i)
		}
	}

	return n
}
