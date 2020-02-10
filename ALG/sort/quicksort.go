package sort

// QuickSort NlgN
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left int, right int) {
	iLeft := left
	iRight := right
	pivot := arr[(left+right)/2]
	for iLeft < iRight {
		for arr[iLeft] < pivot {
			iLeft++
		}
		for arr[iRight] > pivot {
			iRight--
		}
		if iRight <= iLeft {
			break
		}
		arr[iRight], arr[iLeft] = arr[iLeft], arr[iRight]
		if arr[iRight] == pivot {
			iLeft++
		}
		if arr[iLeft] == pivot {
			iRight--
		}
	}
	
	if iLeft == iRight {
		iLeft++
		iRight--
	}
	if iLeft < right {
		quickSort(arr, iLeft, right)
	}
	if iRight > left {
		quickSort(arr, left, iRight)
	}
}
