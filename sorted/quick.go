package sorted

/*
* @author: Chen Chiheng
* @date: 2022/12/26 0026 17:15
* @version: 1.0
* @description: 快速排序
**/

func QuickSort(data []int, left, right int) {
	if left < right {
		index := partition(data, left, right)
		QuickSort(data, left, index)
		QuickSort(data, index+1, right)
	}
}

func partition(data []int, left, right int) int {
	var (
		x = data[right-1]
		i = left - 1
	)
	for j := left; j < right-1; j++ {
		if data[j] <= x {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[right-1] = data[right-1], data[i+1]
	return i + 1
}
