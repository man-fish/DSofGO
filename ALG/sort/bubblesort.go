package sort

func BubbleSort(arr []int) {
	exchange := true
	for j := 1;j < len(arr) && exchange;j++  {
		exchange = false
		for i := 0;i < len(arr)-j ;i++  {
			if arr[i] > arr[i+1] {
				arr[i+1],arr[i] = arr[i],arr[i+1]
				exchange = true
			}
		}
	}
}