package arrs

//题干：

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 0 {
		return 0
	}

	var allNums []int
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			allNums = append(allNums, nums1[i])
			i++
		} else {
			allNums = append(allNums, nums2[j])
			j++
		}
	}

	if i < len(nums1) {
		for ; i < len(nums1); i++ {
			allNums = append(allNums, nums1[i])
		}
	}

	if j < len(nums2) {
		for ; j < len(nums2); j++ {
			allNums = append(allNums, nums2[j])
		}
	}

	index := len(allNums) / 2
	if len(allNums)%2 == 1 {
		return float64(allNums[index])
	} else {
		return float64(allNums[index]+allNums[index-1]) / 2
	}

}

//优化版本：
//思路：因为可以直到中位数的位置，所以在拼allums的过程中找出那个数就OK了

func FindMedianSortedArrays_better(nums1 []int, nums2 []int) float64 {

	length := len(nums1) + len(nums2)
	half := length / 2
	single := false
	if length%2 == 1 {
		single = true
	}

	if length == 0 {
		return 0
	}

	if len(nums1) == 0 {
		if single {
			return float64(nums2[half])
		} else {
			return (float64(nums2[half-1]) + float64(nums2[half])) / 2
		}
	} else if len(nums2) == 0 {
		if single {
			return float64(nums1[half])
		} else {
			return (float64(nums1[half-1]) + float64(nums1[half])) / 2
		}
	}

	i := 0
	j := 0

	oldVal := 0

	for i < len(nums1) && j < len(nums2) {

		if i+j == half && single {
			return float64(getSmaller(nums1[i], nums2[j]))
		} else if i+j-1 == half && !single {
			return (float64(getSmaller(nums1[i], nums2[j])) + float64(oldVal)) / 2
		}

		if nums1[i] <= nums2[j] {
			oldVal = nums1[i]
			i++
		} else {
			nums1[i] = nums2[j]
			j++
		}
	}

	if i == len(nums1) {
		if single {
			return float64(nums2[half-i])
		} else {
			return float64((nums2[half-i] + nums2[half-i-1]) / 2)
		}
	} else {
		if single {
			return float64(nums1[half-j])
		} else {
			return float64((nums1[half-j] + nums1[half-j-1]) / 2)
		}
	}

}

func getSmaller(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
