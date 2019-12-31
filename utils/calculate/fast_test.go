package calculate

import (
	"testing"
)

func Test_fast(t *testing.T)  {
	data := []int{1,48,3,34,10, 23, 43}
	fastSort(data)
	t.Log(data)
}
