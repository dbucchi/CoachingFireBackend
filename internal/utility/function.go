package utility

func Join(elements []string, delimiter string) string {
	if len(elements) == 0 {
		return ""
	}
	result := elements[0]
	for _, element := range elements[1:] {
		result += delimiter + element
	}
	return result
}
