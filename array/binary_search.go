package array

/*
给定一个n个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
写一个函数搜索 nums 中的 target，
如果目标值存在返回下标，
否则返回 -1。
*/

func Binary_Search(nums []int, target int) int {
	return binSearch(0, len(nums)-1, nums, target)
}

func binSearch(start, end int, array []int, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if array[mid] > target {
		return binSearch(start, mid-1, array, target)
	} else if array[mid] < target {
		return binSearch(mid+1, end, array, target)
	} else {
		return mid
	}

}

/*
// 版本二
class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size(); // 定义target在左闭右开的区间里，即：[left, right)
        while (left < right) { // 因为left == right的时候，在[left, right)是无效的空间，所以使用 <
            int middle = left + ((right - left) >> 1);
            if (nums[middle] > target) {
                right = middle; // target 在左区间，在[left, middle)中
            } else if (nums[middle] < target) {
                left = middle + 1; // target 在右区间，在[middle + 1, right)中
            } else { // nums[middle] == target
                return middle; // 数组中找到目标值，直接返回下标
            }
        }
        // 未找到目标值
        return -1;
    }
};

*/
