/*Package gosort 归并排序
1.申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
2.设定两个指针，最初位置分别为两个已经排序序列的起始位置；
3.比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
4.重复步骤 3 直到某一指针达到序列尾；
5.将另一序列剩下的所有元素直接复制到合并序列尾。
*/
package gosort

//MergeSort 归并排序实现
func MergeSort(arr []int) []int {
	mid := len(arr) / 2
	if mid <= 1 {
		return arr
	}
	left := []int{}
	right := []int{}

	// for

	left = MergeSort(left)
	right = MergeSort(right)
	arr = append(left, right...)
	return arr
}
