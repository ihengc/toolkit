package sorted

import "sort"

/*
* @author: Heng ChenChi
* @date: 2022/12/26 0026 10:38
* @version: 1.0
* @description: 插入排序
**/

func InsertionSort(data sort.Interface) {
	for i := 1; i < data.Len(); i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
