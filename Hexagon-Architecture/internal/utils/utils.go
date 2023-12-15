package utils

// RoundUp to round up float.
func RoundUp(v float64) int {
	if v != float64(int(v)) {
		return int(v) + 1
	}
	return int(v)
}
