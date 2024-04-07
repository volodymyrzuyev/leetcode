package FindtheDuplicateNumber

func findDuplicate(nums []int) int {
	gigaMap := make(map[int]byte)
	for _, val := range nums {
		if _, ok := gigaMap[val]; ok {
			return val
		}
		gigaMap[val] = 1
	}
	return 0
}
