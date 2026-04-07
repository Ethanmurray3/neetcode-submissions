func calPoints(operations []string) int {
	scores := []int{}
	for _, operation := range operations {
		if operation == "+" {
			n := len(scores)
			scores = append(scores, scores[n-1] + scores[n-2])
		} else if operation == "D" {
			n := len(scores)
			scores = append(scores, scores[n-1] * 2)
		} else if operation == "C" {
			n := len(scores)
			scores = scores[:n-1]
		} else {
			num, _ := strconv.Atoi(operation)
			scores = append(scores, num)
		}
	}
	var total int
	for _, score := range scores {
		total += score
	}
	return total
}
