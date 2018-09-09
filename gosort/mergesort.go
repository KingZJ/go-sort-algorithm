/*Package gosort 归并排序
1.申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
2.设定两个指针，最初位置分别为两个已经排序序列的起始位置；
3.比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
4.重复步骤 3 直到某一指针达到序列尾；
5.将另一序列剩下的所有元素直接复制到合并序列尾。
*/
package gosort

import (
	"fmt"
)

//MergeSort 归并排序实现
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	left = MergeSort(left)
	right = MergeSort(right)
	leftIndex := 0
	rightIndex := 0
	i := 0
	result := []int{}

	for leftIndex < mid && i < 10 {
		fmt.Println(i, left, right, arr, result, leftIndex, rightIndex)
		if len(right) > 0 && rightIndex < mid && left[leftIndex] > right[rightIndex] {
			result = append(result, right[rightIndex])
			rightIndex++
			i++
			continue
		}
		result = append(result, left[leftIndex])
		leftIndex++
		i++
	}
	arr = append(result, right[rightIndex:]...)
	// fmt.Println(arr, rightIndex)
	return arr
}

func merge(left, right []int) []int {
	result := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
			continue
		}
		result = append(result, left[0])
		left = left[1:]
	}
	result = append(result, append(left, right...)...)
	return result
}

//MergeSort1 归并排序实现
func MergeSort1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	left = MergeSort1(left)
	right = MergeSort1(right)
	arr = merge(left, right)
	return arr
}
