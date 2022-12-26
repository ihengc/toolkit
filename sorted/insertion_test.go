package sorted

import "testing"

/*
* @author: Heng ChenChi
* @date: 2022/12/26 0026 11:39
* @version: 1.0
* @description:
**/

type intList []int

func (i intList) Len() int {
	return len(i)
}

func (i intList) Less(m, n int) bool {
	return i[m] < i[n]
}

func (i intList) Swap(m, n int) {
	i[m], i[n] = i[n], i[m]
}

func TestInsertionSort(t *testing.T) {
	var dList intList = []int{6, 2, 9, 1, 5}
	InsertionSort(dList)
	for i := 1; i < dList.Len(); i++ {
		if dList[i] < dList[i-1] {
			t.Fatal("InsertionSort Error")
		}
	}
}
