package lib

func Max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}
