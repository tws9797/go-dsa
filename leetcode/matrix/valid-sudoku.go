package matrix

// https://leetcode.com/problems/first-bad-version/

func FirstBadVersion(n int) int {

	l, r := 1, n
	first := 0

	for l <= r {
		mid := (l + r) / 2

		if isBadVersion(mid) {

			if mid-1 >= 0 && !isBadVersion(mid-1) {
				return mid
			} else {
				r = mid - 1
			}

		} else {
			l = mid + 1
		}
	}

	return first
}

func FirstBadVersion2(n int) int {

	l, r := 1, n

	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func isBadVersion(version int) bool {
	return false
}
