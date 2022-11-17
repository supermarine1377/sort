package main

import (
	"fmt"

	"github.com/supermarine1377/sort/lib"
)

func main() {
	fmt.Println("test")
}

// O(N^2)
func bubbleSort(arg []int) []int {
	if len(arg) == 0 {
		return nil
	}
	if len(arg) == 1 {
		return arg
	}

	for i := 0; i < len(arg)-1; i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}

	return arg
}

// O(N)のループの中でO(N)で最小値を探しているのでO(N^2)
func selectSort(arg []int) []int {
	for i := 0; i < len(arg); i++ {
		_, minIndex := lib.Min(arg[i:])
		arg[i], arg[i+minIndex] = arg[i+minIndex], arg[i]
	}
	return arg
}
