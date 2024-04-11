package RemoveKDigits

import (
	"fmt"
	"math"
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
	numInt := makeInt(num)

	return fmt.Sprintf("%d", recurse(numInt, k))

}

func recurse(num, k int) int {
	if k == 0 || num == 0 {
		return num
	}
	smalest := num
	numString := fmt.Sprintf("%d", num)
	stringLength := len(numString)

	for pos, val := range numString {
		multFact := stringLength - pos - 1
		valInt := int(val) - 48
		valPosInt := valInt * int(math.Pow10(multFact))
		tempNum := num - valPosInt
		if tempNum < 0 {
			continue
		}
		candidate := recurse(tempNum, k-1)

		if smalest > candidate {
			smalest = candidate
		}
	}

	return smalest
}
