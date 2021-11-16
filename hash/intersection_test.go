package hash

import "testing"

func TestIntersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	res := Intersection(nums1, nums2)

	ShowArrayList(res)
}

func TestIntersection2(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	res := Intersection(nums1, nums2)

	ShowArrayList(res)
}

func ShowArrayList(res []int) {
	for _, val := range res {
		print(val, "\t")
	}
	println()
}
