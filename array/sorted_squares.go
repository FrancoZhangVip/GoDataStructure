package array

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序


进阶：

请你设计时间复杂度为 O(n) 的算法解决本问题

*/

func SortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	minIndex := 0
	minVal := 100000000
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
		if nums[i] < minVal {
			minVal = nums[i]
			minIndex = i
		}
	}

	if minIndex == 0 {
		return nums
	}

	left := minIndex - 1
	right := minIndex
	index := 0
	for left >= 0 && right < len(res) {
		if nums[left] <= nums[right] {

			res[index] = nums[left]
			index++
			left--
		} else {

			res[index] = nums[right]
			index++
			right++
		}
	}
	if left >= 0 {
		for i := left; i >= 0; i-- {

			res[index] = nums[i]
			index++
		}
	}
	if right < len(res) {
		for i := right; i < len(nums); i++ {

			res[index] = nums[i]
			index++
		}
	}
	return res
}

/*

//我们可以使用两个指针分别指向位置 00 和 n-1n−1，每次比较两个指针对应的数，
//选择较大的那个逆序放入答案并移动指针。这种方法无需处理某一指针移动至边界的情况，
//读者可以仔细思考其精髓所在。

func sortedSquares(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    i, j := 0, n-1
    for pos := n - 1; pos >= 0; pos-- {
        if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
            ans[pos] = v
            i++
        } else {
            ans[pos] = w
            j--
        }
    }
    return ans
}


*/
