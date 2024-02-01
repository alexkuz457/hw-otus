package hw03frequencyanalysis

import (
	"fmt"
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

	// мы не можем на этом этапе взять первые 10 элементов слайса.
	// весь слайс мы тоже оставить не можем из соображений производительности (хотя это было бы проще всего)
	// поэтому мы берем первые 10 отличающихся количеств вхождений слова.
	// это значит что если все слова входят в текст по 2 раза нам придется взять весь слайс, но по-другому никак,
	// однако в большинстве случаев слайс после такой частичной обрезки станет меньше и лишние слова сортировать не придется.
	minLen := 0
	for i, k := 1, 0; i < len(sortSlice) && k < 10; i++ {
		if m[sortSlice[i-1]] > m[sortSlice[i]] {
			k++
		}
		minLen = i + 1
	}
	// вот здесь как раз происходит частичная обрезка слайса до величины
	// minLen - минимальной длинны слайса, чтобы можно было его качественно отсортировать
	sortSlice = sortSlice[:minLen]

	// сортируем повторно. Теперь лексикографически (думаю впринципе можно было бы обойтись одной этой сортировкой,
	// на текстах типа Вини-пуха разница 0.1 сек - чтобы почувстововать разницу в качестве текста нужно загонять британскую энциклопедию)
	sort.Slice(sortSlice, func(i, j int) bool {
		var res bool
		if m[sortSlice[i]] == m[sortSlice[j]] {
			res = sortSlice[i] < sortSlice[j]
			fmt.Println(sortSlice[i], "<", sortSlice[j], "=", res)
		} else {
			res = m[sortSlice[i]] > m[sortSlice[j]]
		}
		return res
	})

	if len(sortSlice) > 10 {
		sortSlice = sortSlice[:10]
	}

	return sortSlice
}
