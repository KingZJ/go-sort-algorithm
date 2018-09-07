/*Package gosort 快速排序
1.从数列中挑出一个元素，称为 “基准”（pivot）;
2.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
*/
package gosort

//QuickSort 快速排序实现
func QuickSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	pivot := arr[0]
	left := []int{}
	right := []int{}
	middle := []int{}
	for _, val := range arr {
		if val > pivot {
			right = append(right, val)
		} else if val < pivot {
			left = append(left, val)
		} else {
			middle = append(middle, val)
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	arr = append(left, append(middle, right...)...)
	return arr
}
