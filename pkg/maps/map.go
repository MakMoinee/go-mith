package maps

import (
	"fmt"
	"strings"
)

// GroupMapByNumberInKey() - filter the maps by number in keys. e.g: Test1a,Test2,Test1b. It's expected to have
// map[1:[Test1a Test1b] 2:[Test2]] where Test1a and Test1b are group together because they have common number in keys
// which is 1
func GroupMapByNumberInKey(set map[string]string) map[int][]string {
	groupMap := make(map[int][]string)
	keys := []string{}
	count := 0

	// collate keys
	for key := range set {
		keys = append(keys, key)
	}

	for _, keyText := range keys {
		if len(groupMap) == 0 {
			groupMap[1] = append(groupMap[1], keyText)
			count++
			continue
		} else {
			populateFlag, lastCount := isPopulated(keyText, groupMap)
			if !populateFlag {
				groupMap[lastCount+1] = append(groupMap[lastCount+1], keyText)
			}
			count = lastCount + 1
		}

	}

	return groupMap
}

func isPopulated(key string, group map[int][]string) (bool, int) {
	total := len(group)
	itContains := false
	if total > 0 {
		for i := 0; i < total; i++ {
			contains := strings.Contains(key, fmt.Sprintf("%v", i+1))
			if contains {
				group[i+1] = append(group[i+1], key)
				itContains = contains
				break
			}
		}
	}
	return itContains, total
}
