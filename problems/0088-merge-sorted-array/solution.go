// https://leetcode.com/problems/merge-sorted-array/description/
package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	l, r, i := m-1, n-1, len(nums1)-1

	for ; i >= 0 && l >= 0 && r >= 0; i-- {
		if nums1[l] > nums2[r] {
			nums1[i] = nums1[l]
			l--
		} else {
			nums1[i] = nums2[r]
			r--
		}
	}

	for ; r >= 0; r-- {
		nums1[i] = nums2[r]
		i--
	}
}
