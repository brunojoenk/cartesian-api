package utils

import "strings"

func ConcatStrings(stringsToConcat ...string) string {
	var strBuilder strings.Builder
	for i := 0; i < len(stringsToConcat); i++ {
		strBuilder.WriteString(stringsToConcat[i])
	}
	return strBuilder.String()
}
