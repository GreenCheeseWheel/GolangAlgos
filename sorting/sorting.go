package sorting

func BubbleSort(vec []int) {

	for true {
		was_sorted := true
		for i := 0; i < len(vec)-1; i++ {
			if vec[i+1] < vec[i] {
				helper := vec[i]
				vec[i] = vec[i+1]
				vec[i+1] = helper
				was_sorted = false
			}
		}
		if was_sorted {
			break
		}
	}

}
