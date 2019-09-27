package common

func Sum(numbers []int) int {
	var sum int = 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
