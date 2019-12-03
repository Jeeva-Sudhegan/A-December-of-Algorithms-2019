package main

import "fmt"

func sevenishNumber(n int) int {
	var arr []int
	var sqInd, startInd int = 0, 0
	arr = append(arr, 1)
	for i := 1; i < n; i++ {
		if sqInd != startInd {
			arr = append(arr, arr[sqInd]+arr[startInd])
			startInd++
		} else {
			arr = append(arr, arr[sqInd]*7)
			startInd = 0
			sqInd = i
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Printf("Sevenish number of 10 is %d\n", sevenishNumber(10))
	fmt.Printf("Sevenish number of 5 is %d\n", sevenishNumber(5))
	fmt.Printf("Sevenish number of 15 is %d\n", sevenishNumber(15))
}
