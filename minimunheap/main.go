package main

import (
	"fmt"
)

func heapSort(input []int) {
	inputLen := len(input)
	if inputLen == 0 {
		return
	}

	for i := 0; i < inputLen; i++ {
		minAjust(input[i:])
	}
}

func minAjust(input []int) {
	inputLen := len(input)
	if inputLen <= 1 {
		return
	}

	for i := inputLen/2 - 1; i >= 0; i-- {
		if (2*i+1 <= inputLen-1) && (input[i] >= input[2*i+1]) {
			input[i], input[2*i+1] = input[2*i+1], input[i]
		}

		if (2*i+2 <= inputLen-1) && (input[i] >= input[2*i+2]) {
			input[i], input[2*i+2] = input[2*i+2], input[i]
		}
	}
}

func main() {
	testSlice := []int{2, 3, 1, 4, 5, 6, 2, 1, 23, 43, 1, 32, 3, 5, 2, 1, 8, 54, 4, 0}
	heapSort(testSlice)
	fmt.Println(testSlice)
}
