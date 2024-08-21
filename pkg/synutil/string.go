package synutil

import "strings"

func RuneStr(text string, splitLen int) []string {
	if len(text) <= splitLen {
		return []string{text}
	}
	result := make([]string, 0)
	runes := []rune(text)
	for i := 0; i < len(runes); i += splitLen {
		if i+splitLen < len(runes) {
			result = append(result, string(runes[i:(i+splitLen)]))
		} else {
			result = append(result, string(runes[i:]))
		}
	}
	return result
}

func MakeSpacer(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += " "
	}
	return result
}

func SplitByLineBreak(text string) []string {
	return strings.Split(text, LineBreak)
}
