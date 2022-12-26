package sorted

import "testing"

/*
* @author: Heng ChenChi
* @date: 2022/12/26 0026 14:44
* @version: 1.0
* @description:
**/

func TestMergeSort(t *testing.T) {
	var dList intList = []int{6, 2, 9, 1, 5}
	ret := MergeSort(dList)
	for i := 1; i < len(ret); i++ {
		if ret[i] < ret[i-1] {
			t.Fatal("MergeSort Error")
		}
	}
}
