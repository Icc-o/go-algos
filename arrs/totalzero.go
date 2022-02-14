package arrs

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func threeSum(nums []int) [][]int {
	var mapVlas = make(map[int]int)
	for i, v := range nums {
		mapVlas[v] = i
	}

	var ret [][]int

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if index, ok := mapVlas[0-nums[i]-nums[j]]; ok {
				if index > j {
					ret = append(ret, []int{nums[i], nums[j], nums[index]})
				}
			}
		}
	}

	return ret
}
