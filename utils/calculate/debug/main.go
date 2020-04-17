package main

import (
	"github.com/qingcc/yi/utils/calculate"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	cpuprofile, _ := os.Getwd()
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile + "/heap.prof")
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	arrSource := []int{}
	arr := make([]int, 100000)

	copy(arr, arrSource)
	//log.Println("arr:", arr)
	calculate.QuickSort(arr)

	//copy(arr, arrSource)
	////log.Println("arr:", arr)
	//calculate.HeapSort(arr)

	//
	//copy(arr, arrSource)
	//calculate.ShellSort(arr)
	//
	//copy(arr, arrSource)
	//calculate.InsertSort(arr)
	//
	//copy(arr, arrSource)
	//calculate.SelectSort(arr)
	//
	//copy(arr, arrSource)
	//calculate.MergeSort(arr)
}
