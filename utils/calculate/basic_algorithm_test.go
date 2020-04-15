package calculate

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_bubble(t *testing.T) {
	data := []int{1, 48, 3, 34, 32, 10, 23, 43, 32}
	bubbleSort(data)
	fmt.Println(data)
}

func Test_Select(t *testing.T) {
	data := []int{1, 48, 3, 34, 32, 10, 23, 43, 32}
	SelectSort(data)
	fmt.Println(data)
}
func Test_insert(t *testing.T) {
	data := []int{1, 48, 3, 34, 32, 10, 23, 43, 32}
	InsertSort(data)
	fmt.Println(data)
}

func Test_shell(t *testing.T) {
	data := []int{1, 48, 3, 34, 32, 10, 23, 43, 32}
	ShellSort(data)
	fmt.Println(data)
}
func Test_quick(t *testing.T) {
	data := []int{1, 48, 3, 34, 32, 10, 23, 43, 32}
	QuickSort(data)
	fmt.Println(data)
}

func Test_merge(t *testing.T) {
	data := []int{1, 48, 3, 34, 32, 10, 23, 43, 32}
	Merge_sort(data, 0, len(data)-1)
	fmt.Println(data)
}

func Test_heap(t *testing.T) {
	data := []int{1, 48, 3, 34, 32, 10, 23, 43, 32}
	HeapSort(data)
	fmt.Println(data)
}
func Test_radix(t *testing.T) {
	data := []int{1, 48, 3, 34, 32, 10, 23, 43, 32}
	data = radixSort(data)
	fmt.Println(data)
}

func Test_presure(t *testing.T) {
	var num = 1000000
	var max = 500000
	var input = make([]int, num)
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := 0; i < num; i++ {
		input[i] = r.Intn(max)
	}
	start := time.Now().UnixNano()
	ShellSort(input[:])
	end := time.Now().UnixNano()
	fmt.Println("shell ", len(input), "sort cost: ", end-start, "nano second")

	for i := 0; i < num; i++ {
		input[i] = r.Intn(max)
	}
	start = time.Now().UnixNano()
	QuickSort(input[:])
	end = time.Now().UnixNano()
	fmt.Println("quick ", len(input), "sort cost: ", end-start, "nano second")

	for i := 0; i < num; i++ {
		input[i] = r.Intn(max)
	}
	start = time.Now().UnixNano()
	MergeSort(input[:])
	end = time.Now().UnixNano()
	fmt.Println("merge ", len(input), "sort cost: ", end-start, "nano second")

	for i := 0; i < num; i++ {
		input[i] = r.Intn(max)
	}
	start = time.Now().UnixNano()
	HeapSort(input[:])
	end = time.Now().UnixNano()
	fmt.Println("heap  ", len(input), "sort cost: ", end-start, "nano second")

	//for i := 0; i < num; i++ {
	//	input[i] = r.Intn(max)
	//}
	//start = time.Now().UnixNano()
	//insertSort(input[:])
	//end = time.Now().UnixNano()
	//fmt.Println("insert", len(input), "sort cost: ", end-start, "nano second")

	//for i := 0; i < num; i++ {
	//	input[i] = r.Intn(max)
	//}
	//start = time.Now().UnixNano()
	//selectSort(input[:])
	//end = time.Now().UnixNano()
	//fmt.Println("select", len(input), "sort cost: ", end-start, "nano second")

	//for i := 0; i < num; i++ {
	//	input[i] = r.Intn(max)
	//}
	//start = time.Now().UnixNano()
	//bubbleSort(input[:])
	//end = time.Now().UnixNano()
	//fmt.Println("bubble", len(input), "sort cost: ", end-start, "nano second")

}
