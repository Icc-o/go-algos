package dptr

import (
	"fmt"
	"sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//官方的解体方法排序加双指针
//相当于有三个指针，每一个指针都需要作一下去重。
func ThreeSum(nums []int) [][]int {
	var ret [][]int

	sort.Sort(sort.IntSlice(nums))
	lPtr := 0
	rPtr := 0
	var lastValue int

	fmt.Println(nums)
	for i, _ := range nums {

		if i > 0 && nums[i] == lastValue {
			continue
		}
		lastValue = nums[i]

		if nums[i] > 0 {
			break
		}

		lPtr = i + 1
		rPtr = len(nums) - 1

		for lPtr < rPtr {
			tmp := nums[i] + nums[lPtr] + nums[rPtr]
			if tmp == 0 {
				ret = append(ret, []int{nums[i], nums[lPtr], nums[rPtr]})
				//去重
				for tmp := nums[lPtr]; nums[lPtr] == tmp && lPtr < rPtr; lPtr++ {

				}
				for tmp := nums[rPtr]; nums[rPtr] == tmp && rPtr > lPtr; rPtr-- {

				}
			} else if tmp > 0 {
				rPtr--
			} else {
				lPtr++
			}
		}
	}
	return ret
}
