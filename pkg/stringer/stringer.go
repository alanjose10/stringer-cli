package stringer

func Reverse(input string) string {
	var result string
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

func CountLen(input string) int {
	return len(input)
}
