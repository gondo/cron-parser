package parser

import "strings"

func removeDuplicates(slice []int) (result []int) {
	occurred := map[int]bool{}
	for i := range slice {
		if occurred[slice[i]] != true {
			occurred[slice[i]] = true
			result = append(result, slice[i])
		}
	}
	return result
}

func cleanInput(s string) string {
	return strings.TrimSpace(s)
}

func separateCommand(sections []string, n int) ([]string, string) {
	return sections[:n-1], sections[n-1]
}

func isRange(s string) bool {
	return strings.Contains(s, "-")
}


