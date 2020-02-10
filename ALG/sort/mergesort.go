package sort

// MergeSort 归并排序
func MergeSort(arr, temp []int, start, end int) []int {
	if start < end {
		middle := (start + end) / 2
		MergeSort(arr, temp, start, middle)
		MergeSort(arr, temp, middle+1, end)
		merge(arr, temp, start, middle, end)
	}
	return arr
}

func merge(arr, temp []int, start, mid, end int) {
	i := start
	j := mid + 1
	t := 0
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			t++
			i++
		} else {
			temp[t] = arr[j]
			t++
			j++
		}
	}
	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}

	if j <= end {
		temp[t] = arr[j]
		t++
		j++
	}

	t = 0
	tempstart := start
	for tempstart <= end {
		arr[tempstart] = temp[t]
		tempstart++
		t++
	}
}
