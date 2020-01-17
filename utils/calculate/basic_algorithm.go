package calculate

import (
	"strconv"
)

/*
* 冒泡排序
 */
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return
}

/*
* 选择排序
 */

func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return
}

/*
* 插入排序
 */

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		selected := arr[i] //插入排序是第i个元素和前面i-1个有序数组的比较，交换，比较的一方一直都有元素i
		for j := i - 1; j >= 0 && arr[j] > selected; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return
}

/*
* 希尔排序
 */

func shellSort(arr []int) {
	length := len(arr)
	// 确定h
	h := 1
	for h < length/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < length; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
	//h==1时，相当于插入排序
	//for i:=1; i < length; i++ {
	//	for j:=i; j>=1 && arr[j] < arr[j-1]; j-- {
	//		arr[j], arr[j-1] = arr[j-1], arr[j]
	//	}
	//}
	return
}

/*
* 快速排序
 */

func quickSort(data []int) {
	partition(data, 0, len(data)-1)
}

func partition(data []int, low, high int) {
	if low < high {
		temp := data[(low+high)/2]
		i, j := low, high
		for i <= j {
			for data[i] < temp {
				i++
			}
			for data[j] > temp {
				j--
			}
			if i <= j {
				data[i], data[j] = data[j], data[i]
				i++
				j--
			}
		}
		if low < j {
			partition(data, low, j)
		}
		if i < high {
			partition(data, i, high)
		}
	}
}

/*
* 归并排序
 */

func mergeSort(arr []int) {
	Merge_sort(arr, 0, len(arr)-1)
}

func Merge_sort(arr []int, start int, end int) {
	if start < end {
		middle := (start + end) / 2
		Merge_sort(arr, start, middle)
		Merge_sort(arr, middle+1, end)
		Merge_core(arr, start, middle, end)
	}
}
func Merge_core(arr []int, start int, middle int, end int) {
	var temp []int
	a, b := start, middle+1
	for a <= middle && b <= end {
		if arr[a] < arr[b] {
			temp = append(temp, arr[a])
			a++
		} else {
			temp = append(temp, arr[b])
			b++
		}
	}
	for a <= middle {
		temp = append(temp, arr[a])
		a++
	}
	for b <= end {
		temp = append(temp, arr[b])
		b++
	}
	for i := 0; i < len(temp); i++ {
		arr[start+i] = temp[i]
	}
}

/*
* 堆排序
 */

func heapSort(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		noleaf := (i+1)/2 - 1
		for j := noleaf; j >= 0; j-- {
			HeapAdjust(arr, j, i)
		}
		arr[0], arr[i] = arr[i], arr[0]
	}

}
func HeapAdjust(arr []int, nodeN int, i int) {
	lchild := 2*nodeN + 1
	rchild := lchild + 1
	max := nodeN
	if lchild <= i && arr[lchild] > arr[max] {
		max = lchild
	}
	if rchild <= i && arr[rchild] > arr[max] {
		max = rchild
	}
	if max != nodeN {
		arr[nodeN], arr[max] = arr[max], arr[nodeN]
		HeapAdjust(arr, lchild, i)
		HeapAdjust(arr, rchild, i)
	}
}

/*
* 基础排序
 */
func radixSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	maxl := MaxLen(arr)
	return RadixCore(arr, 0, maxl)
}
func RadixCore(arr []int, digit, maxl int) []int {
	if digit >= maxl {
		return arr
	}
	radix := 10
	count := make([]int, radix)
	bucket := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		count[GetDigit(arr[i], digit)]++
	}
	for i := 1; i < radix; i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		d := GetDigit(arr[i], digit)
		bucket[count[d]-1] = arr[i]
		count[d]--
	}
	return RadixCore(bucket, digit+1, maxl)
}
func GetDigit(x, d int) int {
	a := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	return (x / a[d]) % 10
}
func MaxLen(arr []int) int {
	var maxl, curl int
	for i := 0; i < len(arr); i++ {
		curl = len(strconv.Itoa(arr[i]))
		if curl > maxl {
			maxl = curl
		}
	}
	return maxl
}
