package main

func countArrays(original []int, bounds [][]int) int {
	var count int = 0
	n := len(original)
	var differences []int = make([]int, n)

	for i := 1; i < n-1; i++ {
		differences[i] = original[i] - original[i-1]
	}

	for start := bounds[0][0]; start < bounds[0][1]; start++ {
		isValid := true
		current := start
		for j := 0; j < n; j++ {
			if j > 0 {
				current += differences[j]
			}
			if bounds[j][0] > current || bounds[j][1] < current {
				isValid = false
				break
			}
		}
		if isValid {
			count++
		}
	}
	return count
}

// func main() {
// 	count := countArrays([]int{1, 2, 3, 4}, [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}})
// 	fmt.Println(count)

// }
