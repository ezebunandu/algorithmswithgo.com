package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}
