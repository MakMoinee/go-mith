package maps

import (
	"fmt"
	"strings"
)

var (
	operatorMap = map[string]int{
		">":  1,
		">=": 2,
		"<":  3,
		"<=": 4,
		"==": 5,
		"!=": 6,
	}
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

// FilterNumValueMap() - filter map by number values. Operator indicates the logical operators
// used on the operation (>,>=,==,!=,<,<=). Num is the number criteria for the filtration
func FilterNumValueMap(set map[string]int, operator string, num int) map[string]int {
	finalMap := make(map[string]int)

	if val, exist := operatorMap[operator]; exist {
		filterByOperator(set, val, num)
		finalMap = set
	}
	return finalMap
}

func filterByOperator(set map[string]int, op int, num int) {
	switch op {
	case 1: // >
		for key, val := range set {
			isValid := val > num
			if !isValid {
				delete(set, key)
			}
		}
		break
	case 2:
		for key, val := range set {
			isValid := val >= num
			if !isValid {
				delete(set, key)
			}
		}
		break
	case 3:
		for key, val := range set {
			isValid := val < num
			if !isValid {
				delete(set, key)
			}
		}
		break

	case 4:
		for key, val := range set {
			isValid := val <= num
			if !isValid {
				delete(set, key)
			}
		}
		break

	case 5:
		for key, val := range set {
			isValid := val == num
			if !isValid {
				delete(set, key)
			}
		}
		break
	case 6:
		for key, val := range set {
			isValid := val != num
			if !isValid {
				delete(set, key)
			}
		}
		break
	}
}

// FilterStringValueMapStr() - filters the map with value given. Value property will be filtered
// depending on the operation passed (==,!=). It will be compared to value given in the argument
func FilterStringValueMapStr(set map[string]string, operator string, value string) {
	if op, exist := operatorMap[operator]; exist {
		filterByOperatorStr(set, op, value)
	}
}

func filterByOperatorStr(set map[string]string, op int, chr string) {
	switch op {
	case 5:
		for key, val := range set {
			isValid := val == chr
			if !isValid {
				delete(set, key)
			}
		}
		break
	case 6:
		for key, val := range set {
			isValid := val != chr
			if !isValid {
				delete(set, key)
			}
		}
		break
	}
}
