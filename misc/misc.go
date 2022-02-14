package misc

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/container-with-most-water
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func MaxArea(height []int) int {
	var max = 0
	var tmp int

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			tmp = (j - i) * min(height[i], height[j])
			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}

//初始把双指针指向两侧，可知：
//1、如果把指向短板的指针向内移，可能会变大，
//2、如果把指向长板的指针向内移，一定会变小
func MaxArea_fix(height []int) int {
	var left = 0
	var right = len(height) - 1
	var max = 0
	var tmp = 0

	for left < right {
		tmp = (right - left) * min(height[left], height[right])
		if tmp > max {
			max = tmp
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
