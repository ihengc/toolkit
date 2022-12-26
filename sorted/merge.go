package sorted

/*
* @author: Heng ChenChi
* @date: 2022/12/26 0026 14:13
* @version: 1.0
* @description: 合并排序
**/

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var (
		i, j   = 0, 0
		ll, rl = len(left), len(right)
	)
	ret := make([]int, 0)
	for i < ll && j < rl {
		if left[i] > right[j] {
			ret = append(ret, right[j])
			j++
		} else {
			ret = append(ret, left[i])
			i++
		}
	}
	ret = append(ret, right[j:]...)
	ret = append(ret, left[i:]...)
	return ret
}
