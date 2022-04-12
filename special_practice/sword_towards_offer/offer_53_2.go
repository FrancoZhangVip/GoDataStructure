package sword_towards_offer

func MissingNumber(nums []int) int {
	p, q := 0, len(nums)-1
	for p <= q {
		mid := (p + q) / 2
		if nums[mid] == mid {
			p = mid + 1
		} else {
			q = mid - 1
		}
	}
	return p
}
