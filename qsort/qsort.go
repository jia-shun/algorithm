package qsort

func QuickSort(values []int) []int {

	if len(values) <= 1 {
		return values
	}
	mid := values[0]
	head, tail := 0, len(values)-1
	for i := 1; i <= tail; {
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	QuickSort(values[:head])
	QuickSort(values[head+1:])
	return values

}
