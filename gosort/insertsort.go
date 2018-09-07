/*Package gosort 插入排序
1.将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
2.从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
*/
package gosort

import "fmt"

//InsertSort 插入排序实现
func InsertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		notsort := arr[i]
		arr = append(arr[:i], arr[i+1:]...) //删除位置 i 的元素
		for j := i - 1; j >= 0; j-- {
			if notsort > arr[j] { //插入j+1位置
				fmt.Println(i, arr[:j+1], notsort, arr[j+1:])
				arr = append(arr[:j+1], append([]int{notsort}, arr[j+1:]...)...)
				break
			}
		}
	}
	return arr
}
