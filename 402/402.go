package RemoveKDigits

import (
	"fmt"
	"math"
	"strings"
)

func makeInt(num string) int {
	retNum := 0

	leng := len(num)

	for pos, val := range num {
		multFact := leng - pos - 1
		gigaVal := int(val) - 48
		lapaVal := gigaVal * int(math.Pow(10, float64(multFact)))
		retNum += lapaVal
	}
	return retNum
}

func removeKdigits(num string, k int) string {
	if k == 0 {
		return strings.TrimLeft(num, "0")
	}

	smalest := makeInt(num)

	for pos := range num {
		curString := num[:pos] + num[pos+1:]
		candidate := makeInt(removeKdigits(curString, k-1))
		if smalest > candidate {
			smalest = candidate
		}

	}
	return fmt.Sprintf("%d", smalest)
}
