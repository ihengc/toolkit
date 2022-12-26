package sorted

import "testing"

/*
* @author: Heng ChenChi
* @date: 2022/12/26 0026 15:51
* @version: 1.0
* @description:
**/

func TestBubbleSort(t *testing.T) {
	var dList intList = []int{6, 2, 9, 1, 5}
	BubbleSort(dList)
	for i := 1; i < dList.Len(); i++ {
		if dList[i] < dList[i-1] {
			t.Fatal("BubbleSort Error")
		}
	}
}
