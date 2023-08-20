package common_function

// return true if any slice contains data
func HasData[T any](slice []T) bool {
	LenOfSlice := len(slice)
	if LenOfSlice > 0 {
		return true
	} else {
		return false
	}
}
