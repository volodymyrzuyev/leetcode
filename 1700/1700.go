package NumberofStudentsUnabletoEatLunch

func countStudents(students []int, sandwiches []int) int {
	stuList := make([]int, 2)
	for _, studentSendwich := range students {
		if studentSendwich == sandwiches[0] {
			sandwiches = sandwiches[1:]
			continue
		}
		stuList[studentSendwich]++
	}

	for {
		if len(sandwiches) == 0 {
			break
		}
		if stuList[sandwiches[0]] == 0 {
			return len(sandwiches)
		}
		stuList[sandwiches[0]]--
		sandwiches = sandwiches[1:]
	}

	return 0
}
