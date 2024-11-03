package mystrings

func Reverse(s string) string {
	result := ""
	for _, char := range s {
		result = string(char) + result
	}
	return result
}
