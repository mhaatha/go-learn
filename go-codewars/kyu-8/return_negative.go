package kata

func MakeNegative(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return x
	}
	return x - x - x
}
