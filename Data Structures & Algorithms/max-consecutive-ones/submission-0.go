func findMaxConsecutiveOnes(nums []int) int {
	highestcount := 0
	counter := 0
	for _, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter = 0
		}
		if counter > highestcount {
			highestcount = counter
		}
	}
	return highestcount
}
