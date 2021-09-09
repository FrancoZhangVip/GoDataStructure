package algorithm_init

import (
	"testing"
)

/*
示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5


示例2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
*/

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	res := FindKthLargest(nums, k)
	println("res is : ", res)
	//assert.Equal(t,5,res)

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	res = FindKthLargest(nums, k)
	println("res is : ", res)
	//assert.Equal(t,4,res)
}
