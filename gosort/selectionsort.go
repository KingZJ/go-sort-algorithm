/*Package gosort 选择排序
1.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
3.重复第二步，直到所有元素均排序完毕。
*/
package gosort

//SelectionSort 选择排序实现
func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//SelectionIndexSort 选择排序实现 将最小数|最大数的位置找到与头|尾交换
func SelectionIndexSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i; j < length; j++ {
			if arr[i] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
