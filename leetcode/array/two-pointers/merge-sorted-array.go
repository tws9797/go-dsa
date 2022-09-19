package two_pointers

func Merge(nums1 []int, m int, nums2 []int, n int) {

	for i := m + n - 1; i >= 0; i-- {

		// If n == 0, leave nums1 as it is
		if n == 0 {
			break
		} else if m == 0 {

			// If m == 0, copy the rest of nums2 to nums1
			nums1[i] = nums2[n-1]
			n--
		} else {

			// Put the greater number at the end of nums1
			if nums1[m-1] > nums2[n-1] {
				nums1[i] = nums1[m-1]
				m--
			} else {
				nums1[i] = nums2[n-1]
				n--
			}
		}
	}
}
