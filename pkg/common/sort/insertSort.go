package sort

// InsertSort ..
func InsertSort(data []int) {
	var len int = len(data)
	for i := 1; i < len; i++ {
		var tmp int = data[i]
		j := i - 1
		for ; j >= 0; j-- {
			var aim int = data[j]
			if aim > tmp {
				data[j+1] = data[j]
			} else {
				break
			}
		}
		data[j+1] = tmp
	}
}

// BubbleSort ..
func BubbleSort(data []int) {
	var len int = len(data)
	for i := 0; i < len; i++ {
		var isChange bool
		for j := 0; j < len-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isChange = true
			}
		}
		if false == isChange {
			break
		}
	}
}

// QuikSort ..
func QuikSort(data []int) {

}

func partition(data []int, start, end int) {
	var midNum = data[start]
	start++
	for start >= end {
		for data[start] < midNum && start < end {
			start++
		}

		for data[end] > midNum && start < end {
			end--
		}

		if start < end {
			data[start], data[end] = data[end], data[start]
		}
	}

}
