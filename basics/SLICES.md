# Slices

- [ ] Create new slice
- [x] len
- [x] cap
- [x] full cut slice
- [ ] make
- [ ] copy
- [ ] append

```golang
// Create a backing array: []int{1, 2, 3, 4, 5}
var nums []int = []int{1, 2, 3, 4, 5}

nElemsCopied := copy(nums, []int{20, 30}) // 2
// nums = []int{20, 30, 3, 4, 5}

names := make([]string, 0, 4)
// []string{}, len = 0, cap = 4

names = append(names, "Lemuel") // []string{"Lemuel"}, len = 1, cap = 4

// share the same nums backing array
nums_a = nums[2:4] // []int{3, 4}
len(nums_a) // 2
cap(nums_b) // 3 because []int{3, 4, "5"}

// share the same nums backing array
nums_b = nums[:3] // []int{1, 2, 3}
len(nums_b) // 3
cap(nums_b) // 5, because []int{1, 2, 3, "4", "5"}

// share the same nums backing array
nums_c = nums[3:] // []int{4, 5}
len(nums_c) // 2
cap(nums_c) // 2 


```