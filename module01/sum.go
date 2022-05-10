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

func RecursiveSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else {
		return numbers[0] + RecursiveSum(numbers[1:])
	}
}
