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