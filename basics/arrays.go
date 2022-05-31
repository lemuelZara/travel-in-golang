package basics

var students [][]interface{} = [][]interface{}{
	{"Lemuel", 8, 9, 10},
	{"Maria", 8, 7, 10},
	{"João", 9, 9, 8},
}

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}
	}

	return nums
}

func GetLastArrayElement(arr []int) int {
	last := arr[len(arr)-1:]

	return last[0]
}

func GetAverageGradeByStudentName(name string) float64 {
	var p []interface{}
	var m float64

	for i := 0; i < len(students); i++ {
		if students[i][0] == name {
			p = students[i]

			break
		}
	}

	for j := 0; j < len(p)-1; j++ {
		n, _ := p[j+1].(int)

		m += float64(n)
	}

	return m / float64(len(p[1:]))
}
