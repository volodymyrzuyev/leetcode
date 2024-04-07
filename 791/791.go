package CustomSortString

import "strings"

func customSortString(order string, s string) string {
	var retRune []rune

	sMap := make(map[rune]int)

	for _, val := range s {
		sMap[val]++
	}

	for _, let := range order {
		if count, ok := sMap[let]; ok {
			for i := 0; i < count; i++ {
				retRune = append(retRune, let)
			}
			s = strings.ReplaceAll(s, string(let), "")
		}
	}
	return string(retRune) + s
}
