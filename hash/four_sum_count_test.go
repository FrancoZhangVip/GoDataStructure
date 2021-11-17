package hash

import "testing"

func TestFourSumCount(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}

	println(FourSumCount(nums1, nums2, nums3, nums4))
}

func TestFourSumCount2(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{0}
	nums3 := []int{0}
	nums4 := []int{0}

	println(FourSumCount(nums1, nums2, nums3, nums4))
}

func TestFourSumCount3(t *testing.T) {
	nums1 := []int{-1, -1}
	nums2 := []int{-1, 1}
	nums3 := []int{-1, 1}
	nums4 := []int{1, -1}

	println(FourSumCount(nums1, nums2, nums3, nums4))
}
