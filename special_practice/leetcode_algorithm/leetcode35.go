package leetcode_algorithm

func SearchInsert(nums []int, target int) int {
	p, q := 0, len(nums)
	mid := (p + q) / 2

	for true {
		if p >= q {
			return q
		} else if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			p = mid + 1
			mid = (p + q) / 2
			continue
		} else {
			q = mid
			mid = (p + q) / 2
			continue
		}
	}

	return -1
}
