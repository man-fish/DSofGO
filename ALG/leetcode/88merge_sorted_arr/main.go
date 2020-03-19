/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:
	输入:
		nums1 = [1,2,3,0,0,0], m = 3
		nums2 = [2,5,6],       n = 3

	输出:
		[1,2,2,3,5,6]

*/
package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// reverseIndex := len(nunms1)-1-i

	for i := 0; i < len(nums1); i++ {
		if n == 0 { //如果n=0说明没有必要继续填入数字。
			break
		}

		if m == 0 { //如果说m=0，说明应该把n中的数字填到m中
			nums1[len(nums1)-1-i] = nums2[n-1]
			n--
			continue
			// nums1 = append(nums2[:n], nums1[n:]...)
			//很可惜，不能使用append，因为切片其实不是完全的引用类型，如果改变底层数组映射就不能当作引用类型
			// break
		}

		if i == len(nums1)-1 { //判断是否为最后一位。
			if nums1[0] > nums2[0] {
				nums1[0] = nums2[0]
			}
			break
		}

		if nums1[m-1] > nums2[n-1] { //比较当前m和n的大小
			nums1[len(nums1)-1-i] = nums1[m-1]
			m--
		} else {
			nums1[len(nums1)-1-i] = nums2[n-1]
			n--
		}
	}
}

// nums1 := []int{1, 2, 3, 0, 0, 0}
// nums2 := []int{2, 5, 6}

func main() {
	nums1 := []int{0}
	nums2 := []int{1}

	merge(nums1, 0, nums2, 1)

	fmt.Println(nums1)
}
