package leetcode

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	var res []int
	for k, v := range nums {
		if idx, ok := mp[target-v]; ok {
			res = []int{idx, k}
			break
		}
		mp[v] = k
	}
	return res
}
