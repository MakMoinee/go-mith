package manipulate

import "errors"

// CompareData compares data if its same
func CompareData(data interface{}, desiredDataType interface{}) (bool, error) {
	initialType := 0
	finalType := 0

	switch data.(type) {
	case string:
		initialType = 1
	case int:
		initialType = 2
	case float64:
	case float32:
		initialType = 3
	default:
		initialType = 0
	}

	switch desiredDataType.(type) {
	case string:
		finalType = 1
	case int:
		finalType = 2
	case float64:
	case float32:
		finalType = 3
	default:
		finalType = 0
	}
	if initialType == 0 && finalType == 0 {
		return false, errors.New("data type not being handled")
	}
	res := initialType == finalType
	return res, nil
}
