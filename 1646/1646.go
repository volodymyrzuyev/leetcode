package GetMaximuminGeneratedArray

import "fmt"

func Run() {
	fmt.Println(getMaximumGenerated(4))
}

func getMaximumGenerated(n int) int {
	array := make([]int, n+1)
	array[0] = 0
	array[1] = 1
	max := 0
	for j := 2; j <= n; j++ {
		if j%2 != 0 {
			i := j - 1
			array[j] = array[i/2] + array[(i/2)+1]
			if array[j] > max {
				max = array[j]
			}
			continue
		}
		array[j] = array[j/2]
		if max < array[j] {
			max = array[j]
		}
	}
	return max
}
