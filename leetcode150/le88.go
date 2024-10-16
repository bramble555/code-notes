package leetcode150

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	// 指向 nums1
	i := 0
	// 指向 nums2
	j := 0
	// 指向 nums2
	k := 0
	for i < m && j < n {
		// 小的放到 nums3
		if nums1[i] > nums2[j] {
			nums3[k] = nums2[j]
			j++
		} else {
			nums3[k] = nums1[i]
			i++
		}
		k++
	}
	// 其中有一个 nums 全部放 进去 nums3了,但是不知道哪一个，继续添加另外一个nums
	for i < m {
		nums3[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		nums3[k] = nums2[j]
		j++
		k++
	}
	// 把 nums3 元素放入 nums1
	k = 0
	for ; k < m+n; k++ {
		nums1[k] = nums3[k]
	}
}
