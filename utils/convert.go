package utils

import "strconv"

// Int64ToString Int64 To String
func Int64ToString(d int64) string {
	return strconv.FormatInt(d, 10)
}

func Atoi(s string, default_value int) int {
	if len(s) == 0 {
		return default_value
	}
	if result, err := strconv.Atoi(s); err == nil {
		return result
	} else {
		return default_value
	}
}
