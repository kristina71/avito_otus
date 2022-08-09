package hw09structvalidator

import "strconv"

func stringContains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func intContains(slice []int64, str int64) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func stringsToInts64(strs []string) ([]int64, error) {
	result := make([]int64, 0, len(strs))
	for _, v := range strs {
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}
