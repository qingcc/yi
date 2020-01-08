package calculate

import (
	"fmt"
)


var (
	num int
	name string
	score int
)

func bubble()  {
	fmt.Println("please input students num:")
	fmt.Scanln(&num)
	if num < 1 {
		return
	}

	//input
	arr := make([]int, 0, num)
	for i:=0;i<num; i++ {
		fmt.Println("please input num:")
		fmt.Scanln(&score)
		arr = append(arr, score)
	}

	//bubble sort
	arrBackup := arr
	bubbleSort(arr)
	println(arr)
	selectSort(arrBackup)
	println(arrBackup)
}


/*
* 冒泡排序
 */
func bubbleSort(arr []int)  {
	for i:=0;i<num-1 ; i++ {
		for j := i+1; j < num ; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

/*
* 选择排序
*/

func selectSort(arr []int)  {
	for i:=0 ; i<len(arr)-1 ; i++ {
		max := i
		for j := i+1 ; j < len(arr) ; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[i], arr[max] = arr[max], arr[i]
	}
	fmt.Println("selectSort:", arr)
}

/*
* 插入排序
*/

func insertSort(arr []int) {
	for i:=0;i<len(arr) ; i++ {
		selected := arr[i]
		for j:= i-1; j > 0 && arr[j] > selected ; j-- {
				arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return
}

/*
* 希尔排序
*/

func shellSort(arr []int)  {
	
}

/*
* 快速排序
*/

func fastSort(data []int) {
	partition(data, 0, len(data)-1)
}

func partition(data []int, low, high int) {
	if low < high {
		temp := data[(low+high)/2]
		i, j := low, high
		for i <= j {
			for data[i] > temp {
				i++
			}
			for data[j] < temp {
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
* 合并排序
*/

func mergeSort(arr []int)  {
	
}

/*
* 堆排序
*/

func heapSort(arr []int)  {
	
}

/*
* 基础排序
*/
func radixSort(arr []int)  {
	
}