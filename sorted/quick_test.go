package sorted

import "testing"

/*
* @author: Heng ChenChi
* @date: 2022/12/26 0026 17:17
* @version: 1.0
* @description:
**/

func TestQuickSort(t *testing.T) {
	var dList = []int{6, 2, 9, 1, 5}
	QuickSort(dList, 0, len(dList))
	t.Logf("%v", dList)
	for i := 1; i < len(dList); i++ {
		if dList[i] < dList[i-1] {
			t.Fatal("QuickSort Error")
		}
	}
}
