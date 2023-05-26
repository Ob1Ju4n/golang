package golang

import (
	"fmt"
	"strings"
)

// WordCount - Returns a map of the counts of each word in the
// specified string s
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	countMap := make(map[string]int)

	if len(words) > 0 {
		for _, v := range words {
			countMap[v] = countMap[v] + 1
		}
	}

	fmt.Println(countMap)

	return countMap
}
