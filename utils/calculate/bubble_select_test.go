package calculate

import (
	"fmt"
	"testing"
	"time"
)

func Test_bubble(t *testing.T)  {
	fmt.Println(int64(time.Now().Second()))
	data := []int{1,48,3,34,10, 23, 43}
	bubbleSort(data)
	fmt.Println("bubble:", data)
}

func Test_Select(t *testing.T)  {
	data := []int{1,48,3,34,10, 23, 43}
	selectSort(data)
	fmt.Println(data)
}
func Test_insert (t *testing.T)  {
	data := []int{1,48,3,34, 32,10, 23, 43, 32}
	insertSort(data)
	fmt.Println(data)
}
