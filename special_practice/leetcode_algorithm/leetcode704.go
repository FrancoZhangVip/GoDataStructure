package leetcode_algorithm

func Search(nums []int, target int) int {
	res := -1
	if len(nums) == 0 {
		return res
	}

	p := 0
	q := len(nums)
	mid := (p + q) / 2
	for true {
		if p >= q {
			break
		} else if nums[mid] == target {
			res = mid
			break
		} else if nums[mid] > target {
			q = mid
			mid = (p + q) / 2
			continue
		} else {
			p = mid + 1
			mid = (p + q) / 2
			continue
		}
	}
	return res
}
