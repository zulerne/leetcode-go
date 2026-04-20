// https://leetcode.com/problems/merge-sorted-array/description/
package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	k, l, r := len(nums1)-1, m-1, n-1

	for ; l >= 0 && r >= 0; k-- {
		if nums1[l] > nums2[r] {
			nums1[k] = nums1[l]
			l--
		} else {
			nums1[k] = nums2[r]
			r--
		}
	}

	if r >= 0 {
		for ; k >= 0; k, r = k-1, r-1 {
			nums1[k] = nums2[r]
		}
	}
}
