package iteration

// Repeat func repeats a string specific number of times. returns string
func Repeat(char string, count int) string {
	var result string = ""

	for i := 0; i < count; i++ {
		result += char
	}

	return result
}
