// 42. 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 示例 1：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
// 示例 2：
// 输入：height = [4,2,0,3,2,5]
// 输出：9
package alg

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	var (
		left, right       = 0, len(height) - 1
		maxLeftHeight, maxRightHeight = height[0], height[right]
		area              = 0
	)

	for left < right {
		if maxLeftHeight < maxRightHeight {
			area += maxLeftHeight - height[left]
			left++
			if height[left] > maxLeftHeight {
				maxLeftHeight = height[left]
			}
		} else {
			area += maxRightHeight - height[right]
			right--
			if height[right] > maxRightHeight {
				maxRightHeight = height[right]
			}
		}
	}
	return area
}

// func trap(height []int) int {
// 	if len(height) < 3 {
// 		return 0
// 	}

// 	left, right := 0, len(height)-1
// 	maxLeft := height[left]
// 	maxRight := height[right]
// 	area := 0
// 	for left < right {
// 		// 左边比右g边矮
// 		if maxLeft < maxRight {
// 			area += maxLeft - height[left]
// 			left++
// 			if height[left] > maxLeft {
// 				maxLeft = height[left]
// 			}
// 		} else {
// 			area += maxRight - height[right]
// 			right--
// 			if height[right] > maxRight {
// 				maxRight = height[right]
// 			}
// 		}
// 	}
// 	return area
// }
