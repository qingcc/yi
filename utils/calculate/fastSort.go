package calculate

func fastSort(data []int) {
	partition(data, 0, len(data)-1)
}

func partition(data []int, low, high int) {
	if low < high {
		temp := data[(low+high)/2]
		i, j := low, high
		for i <= j {
			for data[i] > temp {
				i++
			}
			for data[j] < temp {
				j--
			}
			if i <= j {
				data[i], data[j] = data[j], data[i]
				i++
				j--
			}
		}
		if low < j {
			partition(data, low, j)
		}
		if i < high {
			partition(data, i, high)
		}
	}
}
