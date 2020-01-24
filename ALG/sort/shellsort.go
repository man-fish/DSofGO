package sort

func ShellSort(arr []int){
	for dalta := len(arr)/2;dalta > 0;dalta = dalta/2 {
		for i := dalta; i < len(arr) ; i++ {
			temp := arr[i]
			var j int
			for j = i-dalta; j >= 0 && temp < arr[j]; j -= dalta  {
				arr[j+dalta] = arr[j]
			}
			arr[j+dalta] = temp
		}
	}
}