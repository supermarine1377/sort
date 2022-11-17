package lib

// (最小値, 最小値のインデックス)を返す
func Min(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	min := arr[0]
	var index int
	for i, num := range arr {
		if num < min {
			min = num
			index = i
		}
	}
	return min, index
}
