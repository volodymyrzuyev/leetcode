package GoalParserInterpretation

import "fmt"

func Run() {
	fmt.Println(interpret("G()()()()(al)"))
}

func interpret(command string) string {
	var retString string

	for l, r := 0, 0; r <= len(command); r++ {
		if command[l:r] == "G" {
			retString += "G"
			l = r
			continue
		}
		if command[l:r] == "()" {
			retString += "o"
			l = r
			continue
		}
		if command[l:r] == "(al)" {
			retString += "al"
			l = r
			continue
		}
	}
	return retString
}
