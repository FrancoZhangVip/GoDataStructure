package leetcode_algorithm

func SortedSquares(nums []int) []int {
	res := make([]int, 0)
	ans := make([]int, 0)
	if len(nums) == 1 {
		ans = append(ans, nums[0]*nums[0])
		return ans
	}
	for i := 0; i < len(nums); i++ {
		res = append(res, nums[i]*nums[i])
	}

	minIndex := -1
	for i := 0; i < len(res)-1; i++ {
		if res[i] < res[i+1] {
			minIndex = i
			break
		}
	}
	if minIndex == -1 {
		minIndex = len(res) - 1
	}
	//println(minIndex)
	p := minIndex - 1
	q := minIndex + 1
	ans = append(ans, res[minIndex])
	for true {
		if p < 0 && q >= len(res) {
			break
		} else if p < 0 {
			ans = append(ans, res[q])
			q++
		} else if q >= len(res) {
			ans = append(ans, res[p])
			p--
		} else {
			if res[p] <= res[q] {
				ans = append(ans, res[p])
				p--
			} else {
				ans = append(ans, res[q])
				q++
			}
		}
	}

	return ans
}
