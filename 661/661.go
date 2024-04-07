package ImageSmoother

func imageSmoother(img [][]int) [][]int {
	var result [][]int
	m := len(img)
	n := len(img[0])
	for height, row := range img {
		var curRow []int
		for pos := range row {
			sum := 0
			count := 0
			minHeight := max(0, height-1)
			maxHeight := min(height+2, m)
			minWidth := max(0, pos-1)
			maxWidth := min(pos+2, n)
			for _, celRow := range img[minHeight:maxHeight] {
				for _, v := range celRow[minWidth:maxWidth] {
					sum += v
					count += 1
				}
			}
			curRow = append(curRow, sum/count)

		}
		result = append(result, curRow)
	}
	return result
}
