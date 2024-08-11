package kata

func PositiveSum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < 0 {
			continue
		}
		sum += numbers[i]
	}
	return sum
}
