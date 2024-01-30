package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	m := map[string]int{}
	var sortSlice []string
	strAttr := strings.Fields(str)
	for _, word := range strAttr {
		_, ok := m[word]
		if !ok {
			m[word] = 1
			sortSlice = append(sortSlice, word)
		} else {
			m[word]++
		}
	}

	sort.Slice(sortSlice, func(i, j int) bool {
		return m[sortSlice[i]] > m[sortSlice[j]]
	})

	// взять первые 10
	if len(sortSlice) > 10 {
		sortSlice = sortSlice[:10]
	}

	// сортируем повторно. Теперь лексикографически
	sort.Slice(sortSlice, func(i, j int) bool {
		var res bool
		if m[sortSlice[i]] == m[sortSlice[j]] {
			res = sortSlice[i] < sortSlice[j]
		} else {
			res = m[sortSlice[i]] > m[sortSlice[j]]
		}
		return res
	})
	return sortSlice
}
