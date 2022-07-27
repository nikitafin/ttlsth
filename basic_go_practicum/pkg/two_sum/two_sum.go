package main

func TwoSum(nums []int, target int) []int {
	elems := make(map[int]int, len(nums))

	for i, v := range nums {
		if elemIndex, ok := elems[target-v]; ok {
			return []int{i, elemIndex}
		}

		elems[v] = i
	}

	return []int{}
}
