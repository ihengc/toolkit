package sorted

import "sort"

/*
* @author: Heng ChenChi
* @date: 2022/12/26 0026 15:49
* @version: 1.0
* @description: 冒泡排序
**/

func BubbleSort(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		for j := data.Len() - 1; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
