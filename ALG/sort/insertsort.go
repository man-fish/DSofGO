package sort

func InsertSort(arr []int) {
	for i := 1;i < len(arr);i++  {
		temp := arr[i]
		/*这里算作一次移动*/
		var j int
		for j = i-1; j >= 0 && temp < arr[j] ;j--  {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
		/*这里算作一次移动*/
	}
}

/*
	最好情况	顺序
	比较 n-1次
	移动 2(n-1)次

	最坏情况	逆序
	第i趟比较i次，移动i+2次

*/