package iteration

func Repeat(character string, totalRepeat int) string {
	var repeated string
	for i := 0; i < totalRepeat; i++ {
		repeated += character
	}
	return repeated
}
