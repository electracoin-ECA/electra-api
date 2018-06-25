package helpers

// Substring returns a string containing the specified part of the given string.
func Substring(s string, startIndex uint8) string {
	return string([]rune(s)[startIndex:len(s)])
}
