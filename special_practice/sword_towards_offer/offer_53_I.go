package sword_towards_offer

func Search(nums []int, target int) int {
	p, q := 0, len(nums)-1
	index := -1
	counter := 0
	for true {
		mid := (p + q) / 2
		if p > q {
			break
		} else if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] > target {
			q = mid - 1
			continue
		} else {
			p = mid + 1
			continue
		}
	}
	if index == -1 {
		return counter
	}

	for i := index; i < len(nums); i++ {
		if nums[i] == target {
			counter++
		}
	}

	for i := index; i >= 0; i-- {
		if nums[i] == target {
			counter++
		}
	}
	counter--
	return counter
}
