package leetcode_algorithm

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}

func FirstBadVersion(n int) int {
	if n <= 1 {
		return n
	}

	p, q := 2, n
	mid := (p + q) / 2
	for true {
		if !isBadVersion(mid-1) && isBadVersion(mid) {
			return mid
		} else if !isBadVersion(mid) {
			p = mid + 1
			mid = (p + q) / 2
			continue
		} else {
			q = mid - 1
			mid = (p + q) / 2
			continue
		}
	}
	return -1
}
