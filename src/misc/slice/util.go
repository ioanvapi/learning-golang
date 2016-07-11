package util

func isInSlice(a interface{}, list []interface{}) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getSliceIndex(a interface{}, list []interface{}) int {
	for i, b := range list {
		if b == a {
			return i
		}
	}
	return -1
}
