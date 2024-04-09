package TimeNeededtoBuyTickets

func timeRequiredToBuy(tickets []int, k int) int {
	ticketsNeeded := tickets[k]
	time := ticketsNeeded
	for _, ticket := range tickets[:k] {
		if ticket < ticketsNeeded {
			time += ticket
		} else {
			time += ticketsNeeded
		}
	}
	for _, ticket := range tickets[k+1:] {
		if ticket < ticketsNeeded {
			time += ticket
		} else {
			time += ticketsNeeded - 1
		}
	}
	return time
}
