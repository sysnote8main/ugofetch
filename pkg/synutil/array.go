package synutil

import "fmt"

func GetOrDefault(arr []string, index int, defaultValue string) string {
	if index >= len(arr) {
		return defaultValue
	}
	return arr[index]
}

func GetMaxListLen(arrA, arrB []string) int {
	return GetMax(len(arrA), len(arrB))
}

func ListJoin(arrA, arrB []string, defaultA, defaultB, jointStr string) []string {
	result := make([]string, 0)
	for i := range GetMaxListLen(arrA, arrB) {
		result = append(result, fmt.Sprintf(
			"%s%s%s",
			GetOrDefault(arrA, i, defaultA),
			jointStr,
			GetOrDefault(arrB, i, defaultB),
		))
	}
	return result
}
